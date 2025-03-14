from plugify.plugin import Plugin
from plugify.pps import s2sdk as s2, polyhook as pl, dyncall as dc
from customrounds.src.customrounds import CustomRounds

from functools import partial
from pathlib import Path
from enum import IntEnum

import copy

class AcquireMethod(IntEnum):
	PickUp = 0
	Buy = 1

class AcquireResult(IntEnum):
	Allowed = 0
	InvalidItem = 1
	AlreadyOwned = 2
	AlreadyPurchased = 3
	ReachedGrenadeTypeLimit = 4
	ReachedGrenadeTotalLimit = 5
	NotAllowedByTeam = 6
	NotAllowedByMap = 7
	NotAllowedByMode = 8
	NotAllowedForPurchase = 9
	NotAllowedByProhibition = 10

class WeaponEnum(IntEnum):
	WEAPON_DEAGLE = 1
	WEAPON_ELITE = 2
	WEAPON_FIVESEVEN = 3
	WEAPON_GLOCK = 4
	WEAPON_AK47 = 7
	WEAPON_AUG = 8
	WEAPON_AWP = 9
	WEAPON_FAMAS = 10
	WEAPON_G3SG1 = 11
	WEAPON_GALILAR = 13
	WEAPON_M249 = 14
	WEAPON_M4A1 = 16
	WEAPON_MAC10 = 17
	WEAPON_P90 = 19
	WEAPON_ZONE_REPULSOR = 20
	WEAPON_MP5SD = 23
	WEAPON_UMP45 = 24
	WEAPON_XM1014 = 25
	WEAPON_BIZON = 26
	WEAPON_MAG7 = 27
	WEAPON_NEGEV = 28
	WEAPON_SAWEDOFF = 29
	WEAPON_TEC9 = 30
	WEAPON_TASER = 31
	WEAPON_HKP2000 = 32
	WEAPON_MP7 = 33
	WEAPON_MP9 = 34
	WEAPON_NOVA = 35
	WEAPON_P250 = 36
	WEAPON_SHIELD = 37
	WEAPON_SCAR20 = 38
	WEAPON_SG556 = 39
	WEAPON_SSG08 = 40
	WEAPON_KNIFEGG = 41
	WEAPON_KNIFE = 42
	WEAPON_FLASHBANG = 43
	WEAPON_HEGRENADE = 44
	WEAPON_SMOKEGRENADE = 45
	WEAPON_MOLOTOV = 46
	WEAPON_DECOY = 47
	WEAPON_INCGRENADE = 48
	WEAPON_C4 = 49
	ITEM_KEVLAR = 50
	ITEM_ASSAULTSUIT = 51
	ITEM_HEAVYASSAULTSUIT = 52
	ITEM_NVG = 54
	ITEM_DEFUSER = 55
	ITEM_CUTTERS = 56
	WEAPON_HEALTHSHOT = 57
	MUSICKIT_DEFAULT = 58
	WEAPON_KNIFE_T = 59
	WEAPON_M4A1_SILENCER = 60
	WEAPON_USP_SILENCER = 61
	RECIPE_TRADE_UP = 62
	WEAPON_CZ75A = 63
	WEAPON_REVOLVER = 64
	REMOVE_KEYCHAIN_TOOL = 65
	WEAPON_TAGRENADE = 68
	WEAPON_FISTS = 69
	WEAPON_BREACHCHARGE = 70
	WEAPON_TABLET = 72
	WEAPON_MELEE = 74
	WEAPON_AXE = 75
	WEAPON_HAMMER = 76
	WEAPON_SPANNER = 78
	WEAPON_KNIFE_GHOST = 80
	WEAPON_FIREBOMB = 81
	WEAPON_DIVERSION = 82
	WEAPON_FRAG_GRENADE = 83
	WEAPON_SNOWBALL = 84
	WEAPON_BUMPMINE = 85

	@classmethod
	def get_defindex_by_name(cls, name: str):
		name = name.upper()
		try:
			return cls[name].value
		except KeyError:
			return None

	@classmethod
	def get_name_by_defindex(cls, value: int):
		for weapon in cls:
			if weapon.value == value:
				return weapon.name.lower()
		return None

class CSWeaponType(IntEnum):
	WEAPONTYPE_KNIFE = 0x0
	WEAPONTYPE_PISTOL = 0x1
	WEAPONTYPE_SUBMACHINEGUN = 0x2
	WEAPONTYPE_RIFLE = 0x3
	WEAPONTYPE_SHOTGUN = 0x4
	WEAPONTYPE_SNIPER_RIFLE = 0x5
	WEAPONTYPE_MACHINEGUN = 0x6
	WEAPONTYPE_C4 = 0x7
	WEAPONTYPE_TASER = 0x8
	WEAPONTYPE_GRENADE = 0x9
	WEAPONTYPE_EQUIPMENT = 0xa
	WEAPONTYPE_STACKABLEITEM = 0xb
	WEAPONTYPE_FISTS = 0xc
	WEAPONTYPE_BREACHCHARGE = 0xd
	WEAPONTYPE_BUMPMINE = 0xe
	WEAPONTYPE_TABLET = 0xf
	WEAPONTYPE_MELEE = 0x10
	WEAPONTYPE_SHIELD = 0x11
	WEAPONTYPE_ZONE_REPULSOR = 0x12
	WEAPONTYPE_UNKNOWN = 0x13

class CSWeaponCategory(IntEnum):
	WEAPONCATEGORY_OTHER = 0x0
	WEAPONCATEGORY_MELEE = 0x1
	WEAPONCATEGORY_SECONDARY = 0x2
	WEAPONCATEGORY_SMG = 0x3
	WEAPONCATEGORY_RIFLE = 0x4
	WEAPONCATEGORY_HEAVY = 0x5
	WEAPONCATEGORY_COUNT = 0x6

class CSWeaponGearSlot(IntEnum):
	GEAR_SLOT_BOOSTS = 11
	GEAR_SLOT_C4 = 4
	GEAR_SLOT_COUNT = 13
	GEAR_SLOT_FIRST = 0
	GEAR_SLOT_GRENADES = 3
	GEAR_SLOT_INVALID = 4294967295
	GEAR_SLOT_KNIFE = 2
	GEAR_SLOT_LAST = 12
	GEAR_SLOT_PISTOL = 1
	GEAR_SLOT_RESERVED_SLOT10 = 9
	GEAR_SLOT_RESERVED_SLOT11 = 10
	GEAR_SLOT_RESERVED_SLOT6 = 5
	GEAR_SLOT_RESERVED_SLOT7 = 6
	GEAR_SLOT_RESERVED_SLOT8 = 7
	GEAR_SLOT_RESERVED_SLOT9 = 8
	GEAR_SLOT_RIFLE = 0
	GEAR_SLOT_UTILITY = 12

class CRWeapons(Plugin):
	def __init__(self):
		self.cr: CustomRounds = CustomRounds()

		self.gameconfig:				int = 0

		self.can_acquire_ptr:			int = 0
		self.on_weapon_equip_ptr:		int = 0

		self.gamerules:					int = 0
		self.roundtime:					int	= 0

		self.weapons_list:				list = [[] for _ in range(2)]
		self.weapons_storage:			list = [[] for _ in range(64)]
		self.weapons_attributes:		list = [{} for _ in range(2)]

		self.weapons_block:				bool = False
		self.weapons_save:				bool = False
		self.weapons_no_knife:			bool = False
		self.weapons_clear_map:			int  = 0
		self.weapons_no_equip_clear:	int  = 0

		self.vm = None
		self.cs_weapon_data = None

	def plugin_start(self):
		main_plugins_dir = Path(__file__).resolve().parent.parent.parent
		config_path = str(main_plugins_dir / "s2sdk" / "gamedata" / "s2sdk.games.txt")

		self.gameconfig = s2.LoadGameConfigFile(config_path)
		self.cs_weapon_data = s2.GetGameConfigMemSig(self.gameconfig, "GetCSWeaponDataFromKey")

		self.hook_can_acquire()
		self.hook_on_weapon_equip()

		self.cr.on_round_start_register(partial(self.on_round_start))
		self.cr.on_round_end_register(partial(self.on_round_end))
		self.cr.on_player_spawn_register(partial(self.on_player_spawn))

		s2.OnClientPutInServer_Register(partial(self.on_client_put_in_server))
		s2.OnClientDisconnect_Register(partial(self.on_client_disconnect))

		s2.OnServerActivate_Register(partial(self.on_server_activate))

		self.vm = dc.NewVM(4096)
		dc.Mode(self.vm, 0)

	def plugin_end(self):
		dc.Free(self.vm)
		pl.UnhookDetour(self.can_acquire_ptr)
		pl.UnhookDetour(self.on_weapon_equip_ptr)
		del self.can_acquire_ptr
		del self.on_weapon_equip_ptr

	def on_server_activate(self):
		self.gamerules = s2.GetGameRules()
		self.roundtime = s2.GetEntSchema2(self.gamerules, "CCSGameRules", "m_iRoundTime", 0)

	def on_client_put_in_server(self, client):
		if self.weapons_storage[client]:
			self.weapons_storage[client].clear()

	def on_client_disconnect(self, client):
		self.weapons_storage[client].clear()

	def on_player_spawn(self, client):
		if s2.IsClientInGame(client) and s2.IsClientAlive(client):
			if self.cr.current_round_are_custom():

				if self.weapons_save and self.weapons_no_equip_clear < 2:
					self.save_weapons(client)

				if self.weapons_no_equip_clear < 2:
					self.clear_weapons(client)

				self.give_weapons(client)

				if not self.weapons_no_knife and (not self.weapons_list[0] or not self.weapons_list[1]):
					CRWeapons.switch_weapon_on_knife(client)
			else:
				if self.weapons_storage[client]:
					self.clear_weapons(client)
					self.give_weapons(client, True)

		return 0

	def get_weapon_vdata_from_key(self, weapon_name: str = "") -> int:
		dc.Reset(self.vm)
		dc.ArgInt32(self.vm, -1)
		dc.ArgString(self.vm, str(WeaponEnum.get_defindex_by_name(weapon_name)))

		return dc.CallPointer(self.vm, self.cs_weapon_data)

	def set_weapon_attributes(self, default = False):
		for weapon, attributes in self.weapons_attributes[default].items():
			# weapon_awp, {'m_nZoomLevels': 0, 'm_nZoomFOV1': 0, 'm_flZoomTime0': 0.0}
			weapon_vdata = self.get_weapon_vdata_from_key(weapon)

			for attr_key, attr_value in attributes.items():
				if not default:
					if weapon not in self.weapons_attributes[1]:
						self.weapons_attributes[1][weapon] = {}

					if "fl" in attr_key:
						self.weapons_attributes[1][weapon][attr_key] = s2.GetEntSchemaFloat2(weapon_vdata, "CCSWeaponBaseVData", attr_key, 0)
					else:
						self.weapons_attributes[1][weapon][attr_key] = s2.GetEntSchema2(weapon_vdata, "CCSWeaponBaseVData", attr_key, 0)

				if type(attr_value) is float:
					s2.SetEntSchemaFloat2(weapon_vdata, "CCSWeaponBaseVData", attr_key, float(attr_value), True, 0)
				elif type(attr_value) is str:
					s2.SetEntSchemaString2(weapon_vdata, "CCSWeaponBaseVData", attr_key, str(attr_value), True, 0)
				else:
					s2.SetEntSchema2(weapon_vdata, "CCSWeaponBaseVData", attr_key, int(attr_value), True, 0)

	def on_round_start(self):
		if self.cr.current_round_are_custom():
			for i in range(2):
				w_list = self.cr.get_round_setting_dict_value( "weapons", "give" if i == 0 else "ignore", [] )
				if w_list:
					if isinstance(w_list, list):
						self.weapons_list[i] = list(w_list)

			self.weapons_attributes[0]	= copy.deepcopy(self.cr.get_round_setting_dict_value("weapons","attributes", []))
			self.set_weapon_attributes()

			self.weapons_block			= bool(self.cr.get_round_setting_dict_value("weapons","block", 0))
			self.weapons_save			= bool(self.cr.get_round_setting_dict_value("weapons","save", 1))
			self.weapons_no_knife		= bool(self.cr.get_round_setting_dict_value("weapons","no_knife", 0))

			if not self.weapons_no_knife:
				self.weapons_list[1].append("weapon_knife")

			self.weapons_clear_map		= int(self.cr.get_round_setting_dict_value("weapons","clear_map", 0))
			self.weapons_no_equip_clear	= int(self.cr.get_round_setting_dict_value("weapons","no_equip_clear", 0))

			if self.weapons_clear_map == 1 or self.weapons_clear_map == 3:
				CRWeapons.clear_map()

		# self.cr.print(f"\n\t[CR WEAPONS]\n\n\tRound Start\n\n\tCR Status: {is_cr}\n\tWeapons List: {self.weapons_list[0]}\n")

		return 0

	def on_round_end(self):
		if self.cr.current_round_are_custom():
			self.weapons_list[0].clear()
			self.weapons_list[1].clear()

			self.set_weapon_attributes(True)

			self.weapons_attributes[0].clear()
			self.weapons_attributes[1].clear()

			for client in range(64):
				if s2.IsClientInGame(client) and s2.IsClientAlive(client):
					if self.weapons_no_equip_clear == 0 or self.weapons_no_equip_clear == 2:
						self.clear_weapons(client)

					if self.weapons_no_knife:
						weapon_knife = s2.GiveNamedItem(client, "weapon_knife")
						CRWeapons.switch_weapon_on_knife(client, False, weapon_knife)
						continue

					CRWeapons.switch_weapon_on_knife(client)

			if self.weapons_clear_map > 1:
				CRWeapons.clear_map()

			if self.weapons_block:
				self.weapons_block = False

		# self.cr.print(f"\n\t[CR WEAPONS]\n\n\tRound End\n\n\tCR Status: {is_cr}\n\tWeapons List: {self.weapons_list[0]}\n")

		return 0

	def give_weapons(self, client, saved = False):
		slot_priority = [
			CSWeaponGearSlot.GEAR_SLOT_FIRST,
			CSWeaponGearSlot.GEAR_SLOT_PISTOL,
			CSWeaponGearSlot.GEAR_SLOT_KNIFE,
		]

		slot_weapons = {}
		weapons_list = self.weapons_storage[client] if saved else self.weapons_list[0]

		for weapon in weapons_list:
			weapon_hndl = s2.GiveNamedItem(client, weapon)
			weapon_slot = CRWeapons.get_weapon_slot(weapon_hndl)
			if weapon_slot != CSWeaponGearSlot.GEAR_SLOT_INVALID:
				slot_weapons[weapon_slot] = weapon_hndl

		if saved:
			self.weapons_storage[client].clear()

		weapon_target = None

		for slot in slot_priority:
			if slot in slot_weapons:
				weapon_target = slot_weapons[slot]
				break

		s2.SwitchWeapon(client, weapon_target)

	def save_weapons(self, client):
		if self.weapons_storage[client]:
			return

		weapons_list = s2.GetClientWeapons(client)
		for weapon in weapons_list:
			is_melee = s2.GetEntSchema2(s2.GetWeaponVData(weapon), "CCSWeaponBaseVData", "m_bMeleeWeapon", 0)
			if is_melee and self.weapons_no_knife:
				continue

			self.weapons_storage[client].append(s2.GetEntityClassname(weapon))
			# self.cr.print(f"Weapon saved {s2.GetEntityClassname(weapon)}[h: {weapon}]")

	def clear_weapons(self, client):
		weapons_list = s2.GetClientWeapons(client)

		for weapon in weapons_list:
			# m_attribute_manager = s2.GetEntSchema(weapon, "CEconEntity", "m_AttributeManager", 0)
			# m_item = s2.GetEntSchema2(m_attribute_manager, "CAttributeContainer", "m_Item", 0)
			# weapon_def = s2.GetEntSchema2(m_item, "CEconItemView", "m_iItemDefinitionIndex", 0)
			# weapon_def = s2.GetWeaponDefIndex(weapon)

			if CRWeapons.is_weapon_knife(weapon):
				if not self.weapons_no_knife or self.cr.check_round_are_ended():
					CRWeapons.switch_weapon_on_knife(client, False, weapon)
					continue

			s2.RemoveWeapon(client, weapon)

		return 0

	@staticmethod
	def clear_map():
		ent = -1

		while True:
			ent = s2.FindEntityByClassname(ent, "weapon_*")
			if ent == -1:
				break

			owner_entity = s2.GetEntSchemaEnt(ent, "CBaseEntity", "m_hOwnerEntity", 0)
			if owner_entity == -1:
				item_def = s2.GetWeaponDefIndex(ent)

				if item_def == WeaponEnum.WEAPON_DEAGLE or item_def == WeaponEnum.WEAPON_AWP:
					s2.RemoveEntity(ent)

	@staticmethod
	def get_weapon_slot(weapon) -> CSWeaponGearSlot:
		weapon_vdata = s2.GetWeaponVData(weapon)
		weapon_gear_slot = CSWeaponGearSlot(s2.GetEntSchema2(weapon_vdata, "CCSWeaponBaseVData", "m_GearSlot", 0))

		return weapon_gear_slot

	@staticmethod
	def switch_weapon_on_knife(client, enable_iteration = True, weapon = 0):
		weapon_knife = None
		if enable_iteration:
			for weapon in s2.GetClientWeapons(client):
				if CRWeapons.is_weapon_knife(weapon):
					weapon_knife = weapon
					break
		else: weapon_knife = weapon

		if weapon_knife is not None:
			s2.SwitchWeapon(client, weapon_knife)

	@staticmethod
	def is_weapon_knife(weapon):
		weapon_vdata = s2.GetWeaponVData(weapon)
		weapon_type = s2.GetEntSchema2(weapon_vdata, "CCSWeaponBaseVData", "m_WeaponType", 0)

		if weapon_type == CSWeaponType.WEAPONTYPE_KNIFE:
			return True
		return False

	def hook_can_acquire(self):
		can_acquire = s2.GetGameConfigMemSig(self.gameconfig, "CCSPlayer_ItemServices_CanAcquire")

		dt = pl.DataType
		# CPlayer_ItemServices::CanAcquire( CEconItemView *pItem, AcquireMethod::Type acquireMethod, uint16 *pLimit )
		self.can_acquire_ptr = pl.HookDetour(can_acquire, dt.Int8, [dt.Pointer, dt.Pointer, dt.Int32, dt.Pointer])

		if not self.can_acquire_ptr:
			return 0

		pl.AddCallback(self.can_acquire_ptr, False, partial(self.on_can_acquire))

	def on_can_acquire(self, _hook: int, _params: int, _count: int, _ret: int, _post: bool):
		item = pl.GetArgumentPointer(_params, 1)
		if not item:
			return pl.ResultType.Ignored

		if pl.GetArgumentInt32(_params, 2) == AcquireMethod.PickUp:
			if self.weapons_block:
				item_definition = s2.GetEntSchema2(item, "CEconItemView", "m_iItemDefinitionIndex", 0)
				weapon = WeaponEnum.get_name_by_defindex(item_definition)

				if (not weapon in self.weapons_list[0]) and (not weapon in self.weapons_list[1]):
					pl.SetReturnInt8(_ret, AcquireResult.NotAllowedByTeam)
					return pl.ResultType.Supercede

		return pl.ResultType.Ignored

	def hook_on_weapon_equip(self):
		on_weapon_equip = s2.GetGameConfigMemSig(self.gameconfig, "CCSPlayer_WeaponServices::Weapon_Equip")

		dt = pl.DataType
		# CCSPlayer_WeaponServices::Weapon_Equip(CCSPlayer_WeaponServices * this, CBasePlayerWeapon * pWeapon)
		self.on_weapon_equip_ptr = pl.HookDetour(on_weapon_equip, dt.Int64, [dt.Pointer, dt.Pointer])

		if not self.on_weapon_equip_ptr:
			return 0

		pl.AddCallback(self.on_weapon_equip_ptr, True, partial(self.on_weapon_equip))

	def on_weapon_equip(self, _hook: int, _params: int, _count: int, _ret: int, _post: bool):
		weapon_p = pl.GetArgumentPointer(_params, 1)
		weapon_hndl = s2.EntPointerToEntHandle(weapon_p)
		weapon_vdata = s2.GetWeaponVData(weapon_hndl)
		weapon_name = s2.GetEntSchemaString2(weapon_vdata, "CCSWeaponBaseVData", "m_szName", 0)

		if self.cr.current_round_are_custom():
			if weapon_name in self.weapons_attributes[0]:
				if "m_nZoomLevels" in self.weapons_attributes[0][weapon_name]:
					if self.weapons_attributes[0][weapon_name]["m_nZoomLevels"] <= 0:
						roundtime = s2.GetEntSchema2(self.gamerules, "CCSGameRules", "m_iRoundTime", 0)
						s2.SetEntSchema(weapon_hndl, "CBasePlayerWeapon", "m_nNextSecondaryAttackTick", (roundtime * 64) + 10, True, 0)

		return pl.ResultType.Ignored
