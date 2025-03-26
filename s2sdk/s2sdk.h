#pragma once

#include <stdlib.h>
#include <stdint.h>
#include <stdbool.h>

typedef struct String { char* data; size_t size; size_t cap; } String;
typedef struct Vector { void* begin; void* end; void* capacity; } Vector;
typedef struct Vector2 { float x, y; } Vector2;
typedef struct Vector3 { float x, y, z; } Vector3;
typedef struct Vector4 { float x, y, z, w; } Vector4;
typedef struct Matrix4x4 { float m[4][4]; } Matrix4x4;
typedef struct Variant {
    union {
        bool boolean;
        char char8;
        wchar_t char16;
        int8_t int8;
        int16_t int16;
        int32_t int32;
        int64_t int64;
        uint8_t uint8;
        uint16_t uint16;
        uint32_t uint32;
        uint64_t uint64;
        void* ptr;
        float flt;
        double dbl;
        String str;
        Vector vec;
        Vector2 vec2;
        Vector3 vec3;
        Vector4 vec4;
    };
    uint8_t current;
} Variant;

extern void* Plugify_GetMethodPtr(const char* methodName);
extern void Plugify_GetMethodPtr2(const char* methodName, void** addressDest);

static int32_t GetPlayerSlotFromEntityPointer(uintptr_t entity) {
	typedef int32_t (*GetPlayerSlotFromEntityPointerFn)(uintptr_t);
	static GetPlayerSlotFromEntityPointerFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("s2sdk.GetPlayerSlotFromEntityPointer", (void**)&__func);
	return __func(entity);
}

static uintptr_t GetClientBaseFromPlayerSlot(int32_t playerSlot) {
	typedef uintptr_t (*GetClientBaseFromPlayerSlotFn)(int32_t);
	static GetClientBaseFromPlayerSlotFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("s2sdk.GetClientBaseFromPlayerSlot", (void**)&__func);
	return __func(playerSlot);
}

static int32_t GetPlayerSlotFromClientBase(uintptr_t client) {
	typedef int32_t (*GetPlayerSlotFromClientBaseFn)(uintptr_t);
	static GetPlayerSlotFromClientBaseFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("s2sdk.GetPlayerSlotFromClientBase", (void**)&__func);
	return __func(client);
}

static String GetClientAuthId(int32_t playerSlot) {
	typedef String (*GetClientAuthIdFn)(int32_t);
	static GetClientAuthIdFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("s2sdk.GetClientAuthId", (void**)&__func);
	return __func(playerSlot);
}

static uint64_t GetClientAccountId(int32_t playerSlot) {
	typedef uint64_t (*GetClientAccountIdFn)(int32_t);
	static GetClientAccountIdFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("s2sdk.GetClientAccountId", (void**)&__func);
	return __func(playerSlot);
}

static String GetClientIp(int32_t playerSlot) {
	typedef String (*GetClientIpFn)(int32_t);
	static GetClientIpFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("s2sdk.GetClientIp", (void**)&__func);
	return __func(playerSlot);
}

static String GetClientName(int32_t playerSlot) {
	typedef String (*GetClientNameFn)(int32_t);
	static GetClientNameFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("s2sdk.GetClientName", (void**)&__func);
	return __func(playerSlot);
}

static float GetClientTime(int32_t playerSlot) {
	typedef float (*GetClientTimeFn)(int32_t);
	static GetClientTimeFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("s2sdk.GetClientTime", (void**)&__func);
	return __func(playerSlot);
}

static float GetClientLatency(int32_t playerSlot) {
	typedef float (*GetClientLatencyFn)(int32_t);
	static GetClientLatencyFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("s2sdk.GetClientLatency", (void**)&__func);
	return __func(playerSlot);
}

static uint64_t GetUserFlagBits(int32_t playerSlot) {
	typedef uint64_t (*GetUserFlagBitsFn)(int32_t);
	static GetUserFlagBitsFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("s2sdk.GetUserFlagBits", (void**)&__func);
	return __func(playerSlot);
}

static void SetUserFlagBits(int32_t playerSlot, uint64_t flags) {
	typedef void (*SetUserFlagBitsFn)(int32_t, uint64_t);
	static SetUserFlagBitsFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("s2sdk.SetUserFlagBits", (void**)&__func);
	__func(playerSlot, flags);
}

static void AddUserFlags(int32_t playerSlot, uint64_t flags) {
	typedef void (*AddUserFlagsFn)(int32_t, uint64_t);
	static AddUserFlagsFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("s2sdk.AddUserFlags", (void**)&__func);
	__func(playerSlot, flags);
}

static void RemoveUserFlags(int32_t playerSlot, uint64_t flags) {
	typedef void (*RemoveUserFlagsFn)(int32_t, uint64_t);
	static RemoveUserFlagsFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("s2sdk.RemoveUserFlags", (void**)&__func);
	__func(playerSlot, flags);
}

static bool IsClientAuthorized(int32_t playerSlot) {
	typedef bool (*IsClientAuthorizedFn)(int32_t);
	static IsClientAuthorizedFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("s2sdk.IsClientAuthorized", (void**)&__func);
	return __func(playerSlot);
}

static bool IsClientConnected(int32_t playerSlot) {
	typedef bool (*IsClientConnectedFn)(int32_t);
	static IsClientConnectedFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("s2sdk.IsClientConnected", (void**)&__func);
	return __func(playerSlot);
}

static bool IsClientInGame(int32_t playerSlot) {
	typedef bool (*IsClientInGameFn)(int32_t);
	static IsClientInGameFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("s2sdk.IsClientInGame", (void**)&__func);
	return __func(playerSlot);
}

static bool IsClientSourceTV(int32_t playerSlot) {
	typedef bool (*IsClientSourceTVFn)(int32_t);
	static IsClientSourceTVFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("s2sdk.IsClientSourceTV", (void**)&__func);
	return __func(playerSlot);
}

static bool IsClientAlive(int32_t playerSlot) {
	typedef bool (*IsClientAliveFn)(int32_t);
	static IsClientAliveFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("s2sdk.IsClientAlive", (void**)&__func);
	return __func(playerSlot);
}

static bool IsFakeClient(int32_t playerSlot) {
	typedef bool (*IsFakeClientFn)(int32_t);
	static IsFakeClientFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("s2sdk.IsFakeClient", (void**)&__func);
	return __func(playerSlot);
}

static int32_t GetClientTeam(int32_t playerSlot) {
	typedef int32_t (*GetClientTeamFn)(int32_t);
	static GetClientTeamFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("s2sdk.GetClientTeam", (void**)&__func);
	return __func(playerSlot);
}

static int32_t GetClientHealth(int32_t playerSlot) {
	typedef int32_t (*GetClientHealthFn)(int32_t);
	static GetClientHealthFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("s2sdk.GetClientHealth", (void**)&__func);
	return __func(playerSlot);
}

static int32_t GetClientArmor(int32_t playerSlot) {
	typedef int32_t (*GetClientArmorFn)(int32_t);
	static GetClientArmorFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("s2sdk.GetClientArmor", (void**)&__func);
	return __func(playerSlot);
}

static Vector3 GetClientAbsOrigin(int32_t playerSlot) {
	typedef Vector3 (*GetClientAbsOriginFn)(int32_t);
	static GetClientAbsOriginFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("s2sdk.GetClientAbsOrigin", (void**)&__func);
	return __func(playerSlot);
}

static Vector3 GetClientAbsAngles(int32_t playerSlot) {
	typedef Vector3 (*GetClientAbsAnglesFn)(int32_t);
	static GetClientAbsAnglesFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("s2sdk.GetClientAbsAngles", (void**)&__func);
	return __func(playerSlot);
}

static Vector3 GetClientEyeAngles(int32_t playerSlot) {
	typedef Vector3 (*GetClientEyeAnglesFn)(int32_t);
	static GetClientEyeAnglesFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("s2sdk.GetClientEyeAngles", (void**)&__func);
	return __func(playerSlot);
}

static Vector ProcessTargetString(int32_t caller, String* target) {
	typedef Vector (*ProcessTargetStringFn)(int32_t, String*);
	static ProcessTargetStringFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("s2sdk.ProcessTargetString", (void**)&__func);
	return __func(caller, target);
}

static void ChangeClientTeam(int32_t playerSlot, int32_t team) {
	typedef void (*ChangeClientTeamFn)(int32_t, int32_t);
	static ChangeClientTeamFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("s2sdk.ChangeClientTeam", (void**)&__func);
	__func(playerSlot, team);
}

static void SwitchClientTeam(int32_t playerSlot, int32_t team) {
	typedef void (*SwitchClientTeamFn)(int32_t, int32_t);
	static SwitchClientTeamFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("s2sdk.SwitchClientTeam", (void**)&__func);
	__func(playerSlot, team);
}

static void RespawnClient(int32_t playerSlot) {
	typedef void (*RespawnClientFn)(int32_t);
	static RespawnClientFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("s2sdk.RespawnClient", (void**)&__func);
	__func(playerSlot);
}

static void ForcePlayerSuicide(int32_t playerSlot, bool explode, bool force) {
	typedef void (*ForcePlayerSuicideFn)(int32_t, bool, bool);
	static ForcePlayerSuicideFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("s2sdk.ForcePlayerSuicide", (void**)&__func);
	__func(playerSlot, explode, force);
}

static void KickClient(int32_t playerSlot) {
	typedef void (*KickClientFn)(int32_t);
	static KickClientFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("s2sdk.KickClient", (void**)&__func);
	__func(playerSlot);
}

static void BanClient(int32_t playerSlot, float duration, bool kick) {
	typedef void (*BanClientFn)(int32_t, float, bool);
	static BanClientFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("s2sdk.BanClient", (void**)&__func);
	__func(playerSlot, duration, kick);
}

static void BanIdentity(uint64_t steamId, float duration, bool kick) {
	typedef void (*BanIdentityFn)(uint64_t, float, bool);
	static BanIdentityFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("s2sdk.BanIdentity", (void**)&__func);
	__func(steamId, duration, kick);
}

static int32_t GetClientActiveWeapon(int32_t playerSlot) {
	typedef int32_t (*GetClientActiveWeaponFn)(int32_t);
	static GetClientActiveWeaponFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("s2sdk.GetClientActiveWeapon", (void**)&__func);
	return __func(playerSlot);
}

static Vector GetClientWeapons(int32_t playerSlot) {
	typedef Vector (*GetClientWeaponsFn)(int32_t);
	static GetClientWeaponsFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("s2sdk.GetClientWeapons", (void**)&__func);
	return __func(playerSlot);
}

static void DropWeapon(int32_t playerSlot, int32_t weaponHandle, Vector3* target, Vector3* velocity) {
	typedef void (*DropWeaponFn)(int32_t, int32_t, Vector3*, Vector3*);
	static DropWeaponFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("s2sdk.DropWeapon", (void**)&__func);
	__func(playerSlot, weaponHandle, target, velocity);
}

static void StripWeapons(int32_t playerSlot, bool removeSuit) {
	typedef void (*StripWeaponsFn)(int32_t, bool);
	static StripWeaponsFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("s2sdk.StripWeapons", (void**)&__func);
	__func(playerSlot, removeSuit);
}

static void BumpWeapon(int32_t playerSlot, int32_t weaponHandle) {
	typedef void (*BumpWeaponFn)(int32_t, int32_t);
	static BumpWeaponFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("s2sdk.BumpWeapon", (void**)&__func);
	__func(playerSlot, weaponHandle);
}

static void SwitchWeapon(int32_t playerSlot, int32_t weaponHandle) {
	typedef void (*SwitchWeaponFn)(int32_t, int32_t);
	static SwitchWeaponFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("s2sdk.SwitchWeapon", (void**)&__func);
	__func(playerSlot, weaponHandle);
}

static void RemoveWeapon(int32_t playerSlot, int32_t weaponHandle) {
	typedef void (*RemoveWeaponFn)(int32_t, int32_t);
	static RemoveWeaponFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("s2sdk.RemoveWeapon", (void**)&__func);
	__func(playerSlot, weaponHandle);
}

static int32_t GiveNamedItem(int32_t playerSlot, String* itemName) {
	typedef int32_t (*GiveNamedItemFn)(int32_t, String*);
	static GiveNamedItemFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("s2sdk.GiveNamedItem", (void**)&__func);
	return __func(playerSlot, itemName);
}

static uint64_t GetClientButtons(int32_t playerSlot, int32_t buttonIndex) {
	typedef uint64_t (*GetClientButtonsFn)(int32_t, int32_t);
	static GetClientButtonsFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("s2sdk.GetClientButtons", (void**)&__func);
	return __func(playerSlot, buttonIndex);
}

static int32_t GetClientMoney(int32_t playerSlot) {
	typedef int32_t (*GetClientMoneyFn)(int32_t);
	static GetClientMoneyFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("s2sdk.GetClientMoney", (void**)&__func);
	return __func(playerSlot);
}

static void SetClientMoney(int32_t playerSlot, int32_t money) {
	typedef void (*SetClientMoneyFn)(int32_t, int32_t);
	static SetClientMoneyFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("s2sdk.SetClientMoney", (void**)&__func);
	__func(playerSlot, money);
}

static int32_t GetClientKills(int32_t playerSlot) {
	typedef int32_t (*GetClientKillsFn)(int32_t);
	static GetClientKillsFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("s2sdk.GetClientKills", (void**)&__func);
	return __func(playerSlot);
}

static void SetClientKills(int32_t playerSlot, int32_t kills) {
	typedef void (*SetClientKillsFn)(int32_t, int32_t);
	static SetClientKillsFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("s2sdk.SetClientKills", (void**)&__func);
	__func(playerSlot, kills);
}

static int32_t GetClientDeaths(int32_t playerSlot) {
	typedef int32_t (*GetClientDeathsFn)(int32_t);
	static GetClientDeathsFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("s2sdk.GetClientDeaths", (void**)&__func);
	return __func(playerSlot);
}

static void SetClientDeaths(int32_t playerSlot, int32_t deaths) {
	typedef void (*SetClientDeathsFn)(int32_t, int32_t);
	static SetClientDeathsFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("s2sdk.SetClientDeaths", (void**)&__func);
	__func(playerSlot, deaths);
}

static int32_t GetClientAssists(int32_t playerSlot) {
	typedef int32_t (*GetClientAssistsFn)(int32_t);
	static GetClientAssistsFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("s2sdk.GetClientAssists", (void**)&__func);
	return __func(playerSlot);
}

static void SetClientAssists(int32_t playerSlot, int32_t assists) {
	typedef void (*SetClientAssistsFn)(int32_t, int32_t);
	static SetClientAssistsFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("s2sdk.SetClientAssists", (void**)&__func);
	__func(playerSlot, assists);
}

static int32_t GetClientDamage(int32_t playerSlot) {
	typedef int32_t (*GetClientDamageFn)(int32_t);
	static GetClientDamageFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("s2sdk.GetClientDamage", (void**)&__func);
	return __func(playerSlot);
}

static void SetClientDamage(int32_t playerSlot, int32_t damage) {
	typedef void (*SetClientDamageFn)(int32_t, int32_t);
	static SetClientDamageFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("s2sdk.SetClientDamage", (void**)&__func);
	__func(playerSlot, damage);
}

static bool AddAdminCommand(String* name, int64_t adminFlags, String* description, int64_t flags, void* callback, uint8_t type) {
	typedef bool (*AddAdminCommandFn)(String*, int64_t, String*, int64_t, void*, uint8_t);
	static AddAdminCommandFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("s2sdk.AddAdminCommand", (void**)&__func);
	return __func(name, adminFlags, description, flags, callback, type);
}

static bool AddConsoleCommand(String* name, String* description, int64_t flags, void* callback, uint8_t type) {
	typedef bool (*AddConsoleCommandFn)(String*, String*, int64_t, void*, uint8_t);
	static AddConsoleCommandFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("s2sdk.AddConsoleCommand", (void**)&__func);
	return __func(name, description, flags, callback, type);
}

static bool RemoveCommand(String* name, void* callback) {
	typedef bool (*RemoveCommandFn)(String*, void*);
	static RemoveCommandFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("s2sdk.RemoveCommand", (void**)&__func);
	return __func(name, callback);
}

static bool AddCommandListener(String* name, void* callback, uint8_t type) {
	typedef bool (*AddCommandListenerFn)(String*, void*, uint8_t);
	static AddCommandListenerFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("s2sdk.AddCommandListener", (void**)&__func);
	return __func(name, callback, type);
}

static bool RemoveCommandListener(String* name, void* callback, uint8_t type) {
	typedef bool (*RemoveCommandListenerFn)(String*, void*, uint8_t);
	static RemoveCommandListenerFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("s2sdk.RemoveCommandListener", (void**)&__func);
	return __func(name, callback, type);
}

static void ServerCommand(String* command) {
	typedef void (*ServerCommandFn)(String*);
	static ServerCommandFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("s2sdk.ServerCommand", (void**)&__func);
	__func(command);
}

static String ServerCommandEx(String* command) {
	typedef String (*ServerCommandExFn)(String*);
	static ServerCommandExFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("s2sdk.ServerCommandEx", (void**)&__func);
	return __func(command);
}

static void ClientCommand(int32_t playerSlot, String* command) {
	typedef void (*ClientCommandFn)(int32_t, String*);
	static ClientCommandFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("s2sdk.ClientCommand", (void**)&__func);
	__func(playerSlot, command);
}

static void FakeClientCommand(int32_t playerSlot, String* command) {
	typedef void (*FakeClientCommandFn)(int32_t, String*);
	static FakeClientCommandFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("s2sdk.FakeClientCommand", (void**)&__func);
	__func(playerSlot, command);
}

static void PrintToServer(String* msg) {
	typedef void (*PrintToServerFn)(String*);
	static PrintToServerFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("s2sdk.PrintToServer", (void**)&__func);
	__func(msg);
}

static void PrintToConsole(int32_t playerSlot, String* message) {
	typedef void (*PrintToConsoleFn)(int32_t, String*);
	static PrintToConsoleFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("s2sdk.PrintToConsole", (void**)&__func);
	__func(playerSlot, message);
}

static void PrintToChat(int32_t playerSlot, String* message) {
	typedef void (*PrintToChatFn)(int32_t, String*);
	static PrintToChatFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("s2sdk.PrintToChat", (void**)&__func);
	__func(playerSlot, message);
}

static void PrintCenterText(int32_t playerSlot, String* message) {
	typedef void (*PrintCenterTextFn)(int32_t, String*);
	static PrintCenterTextFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("s2sdk.PrintCenterText", (void**)&__func);
	__func(playerSlot, message);
}

static void PrintAlertText(int32_t playerSlot, String* message) {
	typedef void (*PrintAlertTextFn)(int32_t, String*);
	static PrintAlertTextFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("s2sdk.PrintAlertText", (void**)&__func);
	__func(playerSlot, message);
}

static void PrintCentreHtml(int32_t playerSlot, String* message) {
	typedef void (*PrintCentreHtmlFn)(int32_t, String*);
	static PrintCentreHtmlFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("s2sdk.PrintCentreHtml", (void**)&__func);
	__func(playerSlot, message);
}

static void PrintToConsoleAll(String* message) {
	typedef void (*PrintToConsoleAllFn)(String*);
	static PrintToConsoleAllFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("s2sdk.PrintToConsoleAll", (void**)&__func);
	__func(message);
}

static void PrintToChatAll(String* message) {
	typedef void (*PrintToChatAllFn)(String*);
	static PrintToChatAllFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("s2sdk.PrintToChatAll", (void**)&__func);
	__func(message);
}

static void PrintCenterTextAll(String* message) {
	typedef void (*PrintCenterTextAllFn)(String*);
	static PrintCenterTextAllFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("s2sdk.PrintCenterTextAll", (void**)&__func);
	__func(message);
}

static void PrintAlertTextAll(String* message) {
	typedef void (*PrintAlertTextAllFn)(String*);
	static PrintAlertTextAllFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("s2sdk.PrintAlertTextAll", (void**)&__func);
	__func(message);
}

static void PrintCentreHtmlAll(String* message) {
	typedef void (*PrintCentreHtmlAllFn)(String*);
	static PrintCentreHtmlAllFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("s2sdk.PrintCentreHtmlAll", (void**)&__func);
	__func(message);
}

static void PrintToChatColored(int32_t playerSlot, String* message) {
	typedef void (*PrintToChatColoredFn)(int32_t, String*);
	static PrintToChatColoredFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("s2sdk.PrintToChatColored", (void**)&__func);
	__func(playerSlot, message);
}

static void PrintToChatColoredAll(String* message) {
	typedef void (*PrintToChatColoredAllFn)(String*);
	static PrintToChatColoredAllFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("s2sdk.PrintToChatColoredAll", (void**)&__func);
	__func(message);
}

static uint64_t CreateConVar(String* name, String* defaultValue, String* description, int64_t flags) {
	typedef uint64_t (*CreateConVarFn)(String*, String*, String*, int64_t);
	static CreateConVarFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("s2sdk.CreateConVar", (void**)&__func);
	return __func(name, defaultValue, description, flags);
}

static uint64_t CreateConVarBool(String* name, bool defaultValue, String* description, int64_t flags, bool hasMin, bool min, bool hasMax, bool max) {
	typedef uint64_t (*CreateConVarBoolFn)(String*, bool, String*, int64_t, bool, bool, bool, bool);
	static CreateConVarBoolFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("s2sdk.CreateConVarBool", (void**)&__func);
	return __func(name, defaultValue, description, flags, hasMin, min, hasMax, max);
}

static uint64_t CreateConVarInt16(String* name, int16_t defaultValue, String* description, int64_t flags, bool hasMin, int16_t min, bool hasMax, int16_t max) {
	typedef uint64_t (*CreateConVarInt16Fn)(String*, int16_t, String*, int64_t, bool, int16_t, bool, int16_t);
	static CreateConVarInt16Fn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("s2sdk.CreateConVarInt16", (void**)&__func);
	return __func(name, defaultValue, description, flags, hasMin, min, hasMax, max);
}

static uint64_t CreateConVarUInt16(String* name, uint16_t defaultValue, String* description, int64_t flags, bool hasMin, uint16_t min, bool hasMax, uint16_t max) {
	typedef uint64_t (*CreateConVarUInt16Fn)(String*, uint16_t, String*, int64_t, bool, uint16_t, bool, uint16_t);
	static CreateConVarUInt16Fn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("s2sdk.CreateConVarUInt16", (void**)&__func);
	return __func(name, defaultValue, description, flags, hasMin, min, hasMax, max);
}

static uint64_t CreateConVarInt32(String* name, int32_t defaultValue, String* description, int64_t flags, bool hasMin, int32_t min, bool hasMax, int32_t max) {
	typedef uint64_t (*CreateConVarInt32Fn)(String*, int32_t, String*, int64_t, bool, int32_t, bool, int32_t);
	static CreateConVarInt32Fn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("s2sdk.CreateConVarInt32", (void**)&__func);
	return __func(name, defaultValue, description, flags, hasMin, min, hasMax, max);
}

static uint64_t CreateConVarUInt32(String* name, uint32_t defaultValue, String* description, int64_t flags, bool hasMin, uint32_t min, bool hasMax, uint32_t max) {
	typedef uint64_t (*CreateConVarUInt32Fn)(String*, uint32_t, String*, int64_t, bool, uint32_t, bool, uint32_t);
	static CreateConVarUInt32Fn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("s2sdk.CreateConVarUInt32", (void**)&__func);
	return __func(name, defaultValue, description, flags, hasMin, min, hasMax, max);
}

static uint64_t CreateConVarInt64(String* name, int64_t defaultValue, String* description, int64_t flags, bool hasMin, int64_t min, bool hasMax, int64_t max) {
	typedef uint64_t (*CreateConVarInt64Fn)(String*, int64_t, String*, int64_t, bool, int64_t, bool, int64_t);
	static CreateConVarInt64Fn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("s2sdk.CreateConVarInt64", (void**)&__func);
	return __func(name, defaultValue, description, flags, hasMin, min, hasMax, max);
}

static uint64_t CreateConVarUInt64(String* name, uint64_t defaultValue, String* description, int64_t flags, bool hasMin, uint64_t min, bool hasMax, uint64_t max) {
	typedef uint64_t (*CreateConVarUInt64Fn)(String*, uint64_t, String*, int64_t, bool, uint64_t, bool, uint64_t);
	static CreateConVarUInt64Fn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("s2sdk.CreateConVarUInt64", (void**)&__func);
	return __func(name, defaultValue, description, flags, hasMin, min, hasMax, max);
}

static uint64_t CreateConVarFloat(String* name, float defaultValue, String* description, int64_t flags, bool hasMin, float min, bool hasMax, float max) {
	typedef uint64_t (*CreateConVarFloatFn)(String*, float, String*, int64_t, bool, float, bool, float);
	static CreateConVarFloatFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("s2sdk.CreateConVarFloat", (void**)&__func);
	return __func(name, defaultValue, description, flags, hasMin, min, hasMax, max);
}

static uint64_t CreateConVarDouble(String* name, double defaultValue, String* description, int64_t flags, bool hasMin, double min, bool hasMax, double max) {
	typedef uint64_t (*CreateConVarDoubleFn)(String*, double, String*, int64_t, bool, double, bool, double);
	static CreateConVarDoubleFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("s2sdk.CreateConVarDouble", (void**)&__func);
	return __func(name, defaultValue, description, flags, hasMin, min, hasMax, max);
}

static uint64_t CreateConVarColor(String* name, int32_t defaultValue, String* description, int64_t flags, bool hasMin, int32_t min, bool hasMax, int32_t max) {
	typedef uint64_t (*CreateConVarColorFn)(String*, int32_t, String*, int64_t, bool, int32_t, bool, int32_t);
	static CreateConVarColorFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("s2sdk.CreateConVarColor", (void**)&__func);
	return __func(name, defaultValue, description, flags, hasMin, min, hasMax, max);
}

static uint64_t CreateConVarVector2(String* name, Vector2* defaultValue, String* description, int64_t flags, bool hasMin, Vector2* min, bool hasMax, Vector2* max) {
	typedef uint64_t (*CreateConVarVector2Fn)(String*, Vector2*, String*, int64_t, bool, Vector2*, bool, Vector2*);
	static CreateConVarVector2Fn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("s2sdk.CreateConVarVector2", (void**)&__func);
	return __func(name, defaultValue, description, flags, hasMin, min, hasMax, max);
}

static uint64_t CreateConVarVector3(String* name, Vector3* defaultValue, String* description, int64_t flags, bool hasMin, Vector3* min, bool hasMax, Vector3* max) {
	typedef uint64_t (*CreateConVarVector3Fn)(String*, Vector3*, String*, int64_t, bool, Vector3*, bool, Vector3*);
	static CreateConVarVector3Fn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("s2sdk.CreateConVarVector3", (void**)&__func);
	return __func(name, defaultValue, description, flags, hasMin, min, hasMax, max);
}

static uint64_t CreateConVarVector4(String* name, Vector4* defaultValue, String* description, int64_t flags, bool hasMin, Vector4* min, bool hasMax, Vector4* max) {
	typedef uint64_t (*CreateConVarVector4Fn)(String*, Vector4*, String*, int64_t, bool, Vector4*, bool, Vector4*);
	static CreateConVarVector4Fn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("s2sdk.CreateConVarVector4", (void**)&__func);
	return __func(name, defaultValue, description, flags, hasMin, min, hasMax, max);
}

static uint64_t CreateConVarQAngle(String* name, Vector3* defaultValue, String* description, int64_t flags, bool hasMin, Vector3* min, bool hasMax, Vector3* max) {
	typedef uint64_t (*CreateConVarQAngleFn)(String*, Vector3*, String*, int64_t, bool, Vector3*, bool, Vector3*);
	static CreateConVarQAngleFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("s2sdk.CreateConVarQAngle", (void**)&__func);
	return __func(name, defaultValue, description, flags, hasMin, min, hasMax, max);
}

static uint64_t FindConVar(String* name) {
	typedef uint64_t (*FindConVarFn)(String*);
	static FindConVarFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("s2sdk.FindConVar", (void**)&__func);
	return __func(name);
}

static uint64_t FindConVar2(String* name, int16_t type) {
	typedef uint64_t (*FindConVar2Fn)(String*, int16_t);
	static FindConVar2Fn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("s2sdk.FindConVar2", (void**)&__func);
	return __func(name, type);
}

static void HookConVarChange(String* name, void* callback) {
	typedef void (*HookConVarChangeFn)(String*, void*);
	static HookConVarChangeFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("s2sdk.HookConVarChange", (void**)&__func);
	__func(name, callback);
}

static void UnhookConVarChange(String* name, void* callback) {
	typedef void (*UnhookConVarChangeFn)(String*, void*);
	static UnhookConVarChangeFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("s2sdk.UnhookConVarChange", (void**)&__func);
	__func(name, callback);
}

static bool IsConVarFlagSet(uint64_t conVarHandle, int64_t flag) {
	typedef bool (*IsConVarFlagSetFn)(uint64_t, int64_t);
	static IsConVarFlagSetFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("s2sdk.IsConVarFlagSet", (void**)&__func);
	return __func(conVarHandle, flag);
}

static void AddConVarFlags(uint64_t conVarHandle, int64_t flags) {
	typedef void (*AddConVarFlagsFn)(uint64_t, int64_t);
	static AddConVarFlagsFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("s2sdk.AddConVarFlags", (void**)&__func);
	__func(conVarHandle, flags);
}

static void RemoveConVarFlags(uint64_t conVarHandle, int64_t flags) {
	typedef void (*RemoveConVarFlagsFn)(uint64_t, int64_t);
	static RemoveConVarFlagsFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("s2sdk.RemoveConVarFlags", (void**)&__func);
	__func(conVarHandle, flags);
}

static int64_t GetConVarFlags(uint64_t conVarHandle) {
	typedef int64_t (*GetConVarFlagsFn)(uint64_t);
	static GetConVarFlagsFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("s2sdk.GetConVarFlags", (void**)&__func);
	return __func(conVarHandle);
}

static String GetConVarBounds(uint64_t conVarHandle, bool max) {
	typedef String (*GetConVarBoundsFn)(uint64_t, bool);
	static GetConVarBoundsFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("s2sdk.GetConVarBounds", (void**)&__func);
	return __func(conVarHandle, max);
}

static void SetConVarBounds(uint64_t conVarHandle, bool max, String* value) {
	typedef void (*SetConVarBoundsFn)(uint64_t, bool, String*);
	static SetConVarBoundsFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("s2sdk.SetConVarBounds", (void**)&__func);
	__func(conVarHandle, max, value);
}

static String GetConVarDefault(uint64_t conVarHandle) {
	typedef String (*GetConVarDefaultFn)(uint64_t);
	static GetConVarDefaultFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("s2sdk.GetConVarDefault", (void**)&__func);
	return __func(conVarHandle);
}

static String GetConVarValue(uint64_t conVarHandle) {
	typedef String (*GetConVarValueFn)(uint64_t);
	static GetConVarValueFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("s2sdk.GetConVarValue", (void**)&__func);
	return __func(conVarHandle);
}

static Variant GetConVar(uint64_t conVarHandle) {
	typedef Variant (*GetConVarFn)(uint64_t);
	static GetConVarFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("s2sdk.GetConVar", (void**)&__func);
	return __func(conVarHandle);
}

static bool GetConVarBool(uint64_t conVarHandle) {
	typedef bool (*GetConVarBoolFn)(uint64_t);
	static GetConVarBoolFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("s2sdk.GetConVarBool", (void**)&__func);
	return __func(conVarHandle);
}

static int16_t GetConVarInt16(uint64_t conVarHandle) {
	typedef int16_t (*GetConVarInt16Fn)(uint64_t);
	static GetConVarInt16Fn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("s2sdk.GetConVarInt16", (void**)&__func);
	return __func(conVarHandle);
}

static uint16_t GetConVarUInt16(uint64_t conVarHandle) {
	typedef uint16_t (*GetConVarUInt16Fn)(uint64_t);
	static GetConVarUInt16Fn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("s2sdk.GetConVarUInt16", (void**)&__func);
	return __func(conVarHandle);
}

static int32_t GetConVarInt32(uint64_t conVarHandle) {
	typedef int32_t (*GetConVarInt32Fn)(uint64_t);
	static GetConVarInt32Fn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("s2sdk.GetConVarInt32", (void**)&__func);
	return __func(conVarHandle);
}

static uint32_t GetConVarUInt32(uint64_t conVarHandle) {
	typedef uint32_t (*GetConVarUInt32Fn)(uint64_t);
	static GetConVarUInt32Fn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("s2sdk.GetConVarUInt32", (void**)&__func);
	return __func(conVarHandle);
}

static int64_t GetConVarInt64(uint64_t conVarHandle) {
	typedef int64_t (*GetConVarInt64Fn)(uint64_t);
	static GetConVarInt64Fn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("s2sdk.GetConVarInt64", (void**)&__func);
	return __func(conVarHandle);
}

static uint64_t GetConVarUInt64(uint64_t conVarHandle) {
	typedef uint64_t (*GetConVarUInt64Fn)(uint64_t);
	static GetConVarUInt64Fn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("s2sdk.GetConVarUInt64", (void**)&__func);
	return __func(conVarHandle);
}

static float GetConVarFloat(uint64_t conVarHandle) {
	typedef float (*GetConVarFloatFn)(uint64_t);
	static GetConVarFloatFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("s2sdk.GetConVarFloat", (void**)&__func);
	return __func(conVarHandle);
}

static double GetConVarDouble(uint64_t conVarHandle) {
	typedef double (*GetConVarDoubleFn)(uint64_t);
	static GetConVarDoubleFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("s2sdk.GetConVarDouble", (void**)&__func);
	return __func(conVarHandle);
}

static String GetConVarString(uint64_t conVarHandle) {
	typedef String (*GetConVarStringFn)(uint64_t);
	static GetConVarStringFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("s2sdk.GetConVarString", (void**)&__func);
	return __func(conVarHandle);
}

static int32_t GetConVarColor(uint64_t conVarHandle) {
	typedef int32_t (*GetConVarColorFn)(uint64_t);
	static GetConVarColorFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("s2sdk.GetConVarColor", (void**)&__func);
	return __func(conVarHandle);
}

static Vector2 GetConVarVector2(uint64_t conVarHandle) {
	typedef Vector2 (*GetConVarVector2Fn)(uint64_t);
	static GetConVarVector2Fn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("s2sdk.GetConVarVector2", (void**)&__func);
	return __func(conVarHandle);
}

static Vector3 GetConVarVector(uint64_t conVarHandle) {
	typedef Vector3 (*GetConVarVectorFn)(uint64_t);
	static GetConVarVectorFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("s2sdk.GetConVarVector", (void**)&__func);
	return __func(conVarHandle);
}

static Vector4 GetConVarVector4(uint64_t conVarHandle) {
	typedef Vector4 (*GetConVarVector4Fn)(uint64_t);
	static GetConVarVector4Fn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("s2sdk.GetConVarVector4", (void**)&__func);
	return __func(conVarHandle);
}

static Vector3 GetConVarQAngle(uint64_t conVarHandle) {
	typedef Vector3 (*GetConVarQAngleFn)(uint64_t);
	static GetConVarQAngleFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("s2sdk.GetConVarQAngle", (void**)&__func);
	return __func(conVarHandle);
}

static void SetConVarValue(uint64_t conVarHandle, String* value, bool replicate, bool notify) {
	typedef void (*SetConVarValueFn)(uint64_t, String*, bool, bool);
	static SetConVarValueFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("s2sdk.SetConVarValue", (void**)&__func);
	__func(conVarHandle, value, replicate, notify);
}

static void SetConVar(uint64_t conVarHandle, Variant* value, bool replicate, bool notify) {
	typedef void (*SetConVarFn)(uint64_t, Variant*, bool, bool);
	static SetConVarFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("s2sdk.SetConVar", (void**)&__func);
	__func(conVarHandle, value, replicate, notify);
}

static void SetConVarBool(uint64_t conVarHandle, bool value, bool replicate, bool notify) {
	typedef void (*SetConVarBoolFn)(uint64_t, bool, bool, bool);
	static SetConVarBoolFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("s2sdk.SetConVarBool", (void**)&__func);
	__func(conVarHandle, value, replicate, notify);
}

static void SetConVarInt16(uint64_t conVarHandle, int16_t value, bool replicate, bool notify) {
	typedef void (*SetConVarInt16Fn)(uint64_t, int16_t, bool, bool);
	static SetConVarInt16Fn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("s2sdk.SetConVarInt16", (void**)&__func);
	__func(conVarHandle, value, replicate, notify);
}

static void SetConVarUInt16(uint64_t conVarHandle, uint16_t value, bool replicate, bool notify) {
	typedef void (*SetConVarUInt16Fn)(uint64_t, uint16_t, bool, bool);
	static SetConVarUInt16Fn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("s2sdk.SetConVarUInt16", (void**)&__func);
	__func(conVarHandle, value, replicate, notify);
}

static void SetConVarInt32(uint64_t conVarHandle, int32_t value, bool replicate, bool notify) {
	typedef void (*SetConVarInt32Fn)(uint64_t, int32_t, bool, bool);
	static SetConVarInt32Fn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("s2sdk.SetConVarInt32", (void**)&__func);
	__func(conVarHandle, value, replicate, notify);
}

static void SetConVarUInt32(uint64_t conVarHandle, uint32_t value, bool replicate, bool notify) {
	typedef void (*SetConVarUInt32Fn)(uint64_t, uint32_t, bool, bool);
	static SetConVarUInt32Fn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("s2sdk.SetConVarUInt32", (void**)&__func);
	__func(conVarHandle, value, replicate, notify);
}

static void SetConVarInt64(uint64_t conVarHandle, int64_t value, bool replicate, bool notify) {
	typedef void (*SetConVarInt64Fn)(uint64_t, int64_t, bool, bool);
	static SetConVarInt64Fn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("s2sdk.SetConVarInt64", (void**)&__func);
	__func(conVarHandle, value, replicate, notify);
}

static void SetConVarUInt64(uint64_t conVarHandle, uint64_t value, bool replicate, bool notify) {
	typedef void (*SetConVarUInt64Fn)(uint64_t, uint64_t, bool, bool);
	static SetConVarUInt64Fn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("s2sdk.SetConVarUInt64", (void**)&__func);
	__func(conVarHandle, value, replicate, notify);
}

static void SetConVarFloat(uint64_t conVarHandle, float value, bool replicate, bool notify) {
	typedef void (*SetConVarFloatFn)(uint64_t, float, bool, bool);
	static SetConVarFloatFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("s2sdk.SetConVarFloat", (void**)&__func);
	__func(conVarHandle, value, replicate, notify);
}

static void SetConVarDouble(uint64_t conVarHandle, double value, bool replicate, bool notify) {
	typedef void (*SetConVarDoubleFn)(uint64_t, double, bool, bool);
	static SetConVarDoubleFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("s2sdk.SetConVarDouble", (void**)&__func);
	__func(conVarHandle, value, replicate, notify);
}

static void SetConVarString(uint64_t conVarHandle, String* value, bool replicate, bool notify) {
	typedef void (*SetConVarStringFn)(uint64_t, String*, bool, bool);
	static SetConVarStringFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("s2sdk.SetConVarString", (void**)&__func);
	__func(conVarHandle, value, replicate, notify);
}

static void SetConVarColor(uint64_t conVarHandle, int32_t value, bool replicate, bool notify) {
	typedef void (*SetConVarColorFn)(uint64_t, int32_t, bool, bool);
	static SetConVarColorFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("s2sdk.SetConVarColor", (void**)&__func);
	__func(conVarHandle, value, replicate, notify);
}

static void SetConVarVector2(uint64_t conVarHandle, Vector2* value, bool replicate, bool notify) {
	typedef void (*SetConVarVector2Fn)(uint64_t, Vector2*, bool, bool);
	static SetConVarVector2Fn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("s2sdk.SetConVarVector2", (void**)&__func);
	__func(conVarHandle, value, replicate, notify);
}

static void SetConVarVector3(uint64_t conVarHandle, Vector3* value, bool replicate, bool notify) {
	typedef void (*SetConVarVector3Fn)(uint64_t, Vector3*, bool, bool);
	static SetConVarVector3Fn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("s2sdk.SetConVarVector3", (void**)&__func);
	__func(conVarHandle, value, replicate, notify);
}

static void SetConVarVector4(uint64_t conVarHandle, Vector4* value, bool replicate, bool notify) {
	typedef void (*SetConVarVector4Fn)(uint64_t, Vector4*, bool, bool);
	static SetConVarVector4Fn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("s2sdk.SetConVarVector4", (void**)&__func);
	__func(conVarHandle, value, replicate, notify);
}

static void SetConVarQAngle(uint64_t conVarHandle, Vector3* value, bool replicate, bool notify) {
	typedef void (*SetConVarQAngleFn)(uint64_t, Vector3*, bool, bool);
	static SetConVarQAngleFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("s2sdk.SetConVarQAngle", (void**)&__func);
	__func(conVarHandle, value, replicate, notify);
}

static void SendConVarValue(int32_t playerSlot, uint64_t conVarHandle, String* value) {
	typedef void (*SendConVarValueFn)(int32_t, uint64_t, String*);
	static SendConVarValueFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("s2sdk.SendConVarValue", (void**)&__func);
	__func(playerSlot, conVarHandle, value);
}

static String GetClientConVarValue(int32_t playerSlot, String* convarName) {
	typedef String (*GetClientConVarValueFn)(int32_t, String*);
	static GetClientConVarValueFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("s2sdk.GetClientConVarValue", (void**)&__func);
	return __func(playerSlot, convarName);
}

static void SetFakeClientConVarValue(int32_t playerSlot, String* convarName, String* convarValue) {
	typedef void (*SetFakeClientConVarValueFn)(int32_t, String*, String*);
	static SetFakeClientConVarValueFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("s2sdk.SetFakeClientConVarValue", (void**)&__func);
	__func(playerSlot, convarName, convarValue);
}

static String GetGameDirectory() {
	typedef String (*GetGameDirectoryFn)();
	static GetGameDirectoryFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("s2sdk.GetGameDirectory", (void**)&__func);
	return __func();
}

static String GetCurrentMap() {
	typedef String (*GetCurrentMapFn)();
	static GetCurrentMapFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("s2sdk.GetCurrentMap", (void**)&__func);
	return __func();
}

static bool IsMapValid(String* mapname) {
	typedef bool (*IsMapValidFn)(String*);
	static IsMapValidFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("s2sdk.IsMapValid", (void**)&__func);
	return __func(mapname);
}

static float GetGameTime() {
	typedef float (*GetGameTimeFn)();
	static GetGameTimeFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("s2sdk.GetGameTime", (void**)&__func);
	return __func();
}

static int32_t GetGameTickCount() {
	typedef int32_t (*GetGameTickCountFn)();
	static GetGameTickCountFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("s2sdk.GetGameTickCount", (void**)&__func);
	return __func();
}

static float GetGameFrameTime() {
	typedef float (*GetGameFrameTimeFn)();
	static GetGameFrameTimeFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("s2sdk.GetGameFrameTime", (void**)&__func);
	return __func();
}

static double GetEngineTime() {
	typedef double (*GetEngineTimeFn)();
	static GetEngineTimeFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("s2sdk.GetEngineTime", (void**)&__func);
	return __func();
}

static int32_t GetMaxClients() {
	typedef int32_t (*GetMaxClientsFn)();
	static GetMaxClientsFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("s2sdk.GetMaxClients", (void**)&__func);
	return __func();
}

static int32_t PrecacheGeneric(String* model) {
	typedef int32_t (*PrecacheGenericFn)(String*);
	static PrecacheGenericFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("s2sdk.PrecacheGeneric", (void**)&__func);
	return __func(model);
}

static bool IsGenericPrecache(String* model) {
	typedef bool (*IsGenericPrecacheFn)(String*);
	static IsGenericPrecacheFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("s2sdk.IsGenericPrecache", (void**)&__func);
	return __func(model);
}

static int32_t PrecacheModel(String* model) {
	typedef int32_t (*PrecacheModelFn)(String*);
	static PrecacheModelFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("s2sdk.PrecacheModel", (void**)&__func);
	return __func(model);
}

static bool IsModelPrecache(String* model) {
	typedef bool (*IsModelPrecacheFn)(String*);
	static IsModelPrecacheFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("s2sdk.IsModelPrecache", (void**)&__func);
	return __func(model);
}

static bool PrecacheSound(String* sound, bool preload) {
	typedef bool (*PrecacheSoundFn)(String*, bool);
	static PrecacheSoundFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("s2sdk.PrecacheSound", (void**)&__func);
	return __func(sound, preload);
}

static bool IsSoundPrecached(String* sound) {
	typedef bool (*IsSoundPrecachedFn)(String*);
	static IsSoundPrecachedFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("s2sdk.IsSoundPrecached", (void**)&__func);
	return __func(sound);
}

static int32_t PrecacheDecal(String* decal, bool preload) {
	typedef int32_t (*PrecacheDecalFn)(String*, bool);
	static PrecacheDecalFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("s2sdk.PrecacheDecal", (void**)&__func);
	return __func(decal, preload);
}

static bool IsDecalPrecached(String* decal) {
	typedef bool (*IsDecalPrecachedFn)(String*);
	static IsDecalPrecachedFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("s2sdk.IsDecalPrecached", (void**)&__func);
	return __func(decal);
}

static uintptr_t GetEconItemSystem() {
	typedef uintptr_t (*GetEconItemSystemFn)();
	static GetEconItemSystemFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("s2sdk.GetEconItemSystem", (void**)&__func);
	return __func();
}

static bool IsServerPaused() {
	typedef bool (*IsServerPausedFn)();
	static IsServerPausedFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("s2sdk.IsServerPaused", (void**)&__func);
	return __func();
}

static void QueueTaskForNextFrame(void* callback, Vector* userData) {
	typedef void (*QueueTaskForNextFrameFn)(void*, Vector*);
	static QueueTaskForNextFrameFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("s2sdk.QueueTaskForNextFrame", (void**)&__func);
	__func(callback, userData);
}

static void QueueTaskForNextWorldUpdate(void* callback, Vector* userData) {
	typedef void (*QueueTaskForNextWorldUpdateFn)(void*, Vector*);
	static QueueTaskForNextWorldUpdateFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("s2sdk.QueueTaskForNextWorldUpdate", (void**)&__func);
	__func(callback, userData);
}

static float GetSoundDuration(String* name) {
	typedef float (*GetSoundDurationFn)(String*);
	static GetSoundDurationFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("s2sdk.GetSoundDuration", (void**)&__func);
	return __func(name);
}

static void EmitSound(int32_t entityHandle, String* sound, int32_t pitch, float volume, float delay) {
	typedef void (*EmitSoundFn)(int32_t, String*, int32_t, float, float);
	static EmitSoundFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("s2sdk.EmitSound", (void**)&__func);
	__func(entityHandle, sound, pitch, volume, delay);
}

static void EmitSoundToClient(int32_t playerSlot, int32_t channel, String* sound, float volume, int32_t soundLevel, int32_t flags, int32_t pitch, Vector3* origin, float soundTime) {
	typedef void (*EmitSoundToClientFn)(int32_t, int32_t, String*, float, int32_t, int32_t, int32_t, Vector3*, float);
	static EmitSoundToClientFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("s2sdk.EmitSoundToClient", (void**)&__func);
	__func(playerSlot, channel, sound, volume, soundLevel, flags, pitch, origin, soundTime);
}

static uintptr_t EntIndexToEntPointer(int32_t entityIndex) {
	typedef uintptr_t (*EntIndexToEntPointerFn)(int32_t);
	static EntIndexToEntPointerFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("s2sdk.EntIndexToEntPointer", (void**)&__func);
	return __func(entityIndex);
}

static int32_t EntPointerToEntIndex(uintptr_t entity) {
	typedef int32_t (*EntPointerToEntIndexFn)(uintptr_t);
	static EntPointerToEntIndexFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("s2sdk.EntPointerToEntIndex", (void**)&__func);
	return __func(entity);
}

static int32_t EntPointerToEntHandle(uintptr_t entity) {
	typedef int32_t (*EntPointerToEntHandleFn)(uintptr_t);
	static EntPointerToEntHandleFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("s2sdk.EntPointerToEntHandle", (void**)&__func);
	return __func(entity);
}

static uintptr_t EntHandleToEntPointer(int32_t entityHandle) {
	typedef uintptr_t (*EntHandleToEntPointerFn)(int32_t);
	static EntHandleToEntPointerFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("s2sdk.EntHandleToEntPointer", (void**)&__func);
	return __func(entityHandle);
}

static int32_t EntIndexToEntHandle(int32_t entityIndex) {
	typedef int32_t (*EntIndexToEntHandleFn)(int32_t);
	static EntIndexToEntHandleFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("s2sdk.EntIndexToEntHandle", (void**)&__func);
	return __func(entityIndex);
}

static int32_t EntHandleToEntIndex(int32_t entityHandle) {
	typedef int32_t (*EntHandleToEntIndexFn)(int32_t);
	static EntHandleToEntIndexFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("s2sdk.EntHandleToEntIndex", (void**)&__func);
	return __func(entityHandle);
}

static bool IsValidEntHandle(int32_t entityHandle) {
	typedef bool (*IsValidEntHandleFn)(int32_t);
	static IsValidEntHandleFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("s2sdk.IsValidEntHandle", (void**)&__func);
	return __func(entityHandle);
}

static bool IsValidEntPointer(uintptr_t entity) {
	typedef bool (*IsValidEntPointerFn)(uintptr_t);
	static IsValidEntPointerFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("s2sdk.IsValidEntPointer", (void**)&__func);
	return __func(entity);
}

static uintptr_t GetFirstActiveEntity() {
	typedef uintptr_t (*GetFirstActiveEntityFn)();
	static GetFirstActiveEntityFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("s2sdk.GetFirstActiveEntity", (void**)&__func);
	return __func();
}

static uintptr_t GetConcreteEntityListPointer() {
	typedef uintptr_t (*GetConcreteEntityListPointerFn)();
	static GetConcreteEntityListPointerFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("s2sdk.GetConcreteEntityListPointer", (void**)&__func);
	return __func();
}

static bool HookEntityOutput(String* szClassname, String* szOutput, void* callback, uint8_t type) {
	typedef bool (*HookEntityOutputFn)(String*, String*, void*, uint8_t);
	static HookEntityOutputFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("s2sdk.HookEntityOutput", (void**)&__func);
	return __func(szClassname, szOutput, callback, type);
}

static bool UnhookEntityOutput(String* szClassname, String* szOutput, void* callback, uint8_t type) {
	typedef bool (*UnhookEntityOutputFn)(String*, String*, void*, uint8_t);
	static UnhookEntityOutputFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("s2sdk.UnhookEntityOutput", (void**)&__func);
	return __func(szClassname, szOutput, callback, type);
}

static int32_t FindEntityByClassname(int32_t startEntity, String* classname) {
	typedef int32_t (*FindEntityByClassnameFn)(int32_t, String*);
	static FindEntityByClassnameFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("s2sdk.FindEntityByClassname", (void**)&__func);
	return __func(startEntity, classname);
}

static int32_t FindEntityByName(int32_t startEntity, String* name) {
	typedef int32_t (*FindEntityByNameFn)(int32_t, String*);
	static FindEntityByNameFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("s2sdk.FindEntityByName", (void**)&__func);
	return __func(startEntity, name);
}

static int32_t CreateEntityByName(String* className) {
	typedef int32_t (*CreateEntityByNameFn)(String*);
	static CreateEntityByNameFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("s2sdk.CreateEntityByName", (void**)&__func);
	return __func(className);
}

static void DispatchSpawn(int32_t entityHandle) {
	typedef void (*DispatchSpawnFn)(int32_t);
	static DispatchSpawnFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("s2sdk.DispatchSpawn", (void**)&__func);
	__func(entityHandle);
}

static void DispatchSpawn2(int32_t entityHandle, Vector* keys, Vector* values) {
	typedef void (*DispatchSpawn2Fn)(int32_t, Vector*, Vector*);
	static DispatchSpawn2Fn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("s2sdk.DispatchSpawn2", (void**)&__func);
	__func(entityHandle, keys, values);
}

static void RemoveEntity(int32_t entityHandle) {
	typedef void (*RemoveEntityFn)(int32_t);
	static RemoveEntityFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("s2sdk.RemoveEntity", (void**)&__func);
	__func(entityHandle);
}

static String GetEntityClassname(int32_t entityHandle) {
	typedef String (*GetEntityClassnameFn)(int32_t);
	static GetEntityClassnameFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("s2sdk.GetEntityClassname", (void**)&__func);
	return __func(entityHandle);
}

static String GetEntityName(int32_t entityHandle) {
	typedef String (*GetEntityNameFn)(int32_t);
	static GetEntityNameFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("s2sdk.GetEntityName", (void**)&__func);
	return __func(entityHandle);
}

static void SetEntityName(int32_t entityHandle, String* name) {
	typedef void (*SetEntityNameFn)(int32_t, String*);
	static SetEntityNameFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("s2sdk.SetEntityName", (void**)&__func);
	__func(entityHandle, name);
}

static int32_t GetEntityMoveType(int32_t entityHandle) {
	typedef int32_t (*GetEntityMoveTypeFn)(int32_t);
	static GetEntityMoveTypeFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("s2sdk.GetEntityMoveType", (void**)&__func);
	return __func(entityHandle);
}

static void SetEntityMoveType(int32_t entityHandle, int32_t moveType) {
	typedef void (*SetEntityMoveTypeFn)(int32_t, int32_t);
	static SetEntityMoveTypeFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("s2sdk.SetEntityMoveType", (void**)&__func);
	__func(entityHandle, moveType);
}

static float GetEntityGravity(int32_t entityHandle) {
	typedef float (*GetEntityGravityFn)(int32_t);
	static GetEntityGravityFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("s2sdk.GetEntityGravity", (void**)&__func);
	return __func(entityHandle);
}

static void SetEntityGravity(int32_t entityHandle, float gravity) {
	typedef void (*SetEntityGravityFn)(int32_t, float);
	static SetEntityGravityFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("s2sdk.SetEntityGravity", (void**)&__func);
	__func(entityHandle, gravity);
}

static int32_t GetEntityFlags(int32_t entityHandle) {
	typedef int32_t (*GetEntityFlagsFn)(int32_t);
	static GetEntityFlagsFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("s2sdk.GetEntityFlags", (void**)&__func);
	return __func(entityHandle);
}

static void SetEntityFlags(int32_t entityHandle, int32_t flags) {
	typedef void (*SetEntityFlagsFn)(int32_t, int32_t);
	static SetEntityFlagsFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("s2sdk.SetEntityFlags", (void**)&__func);
	__func(entityHandle, flags);
}

static int32_t GetEntityRenderColor(int32_t entityHandle) {
	typedef int32_t (*GetEntityRenderColorFn)(int32_t);
	static GetEntityRenderColorFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("s2sdk.GetEntityRenderColor", (void**)&__func);
	return __func(entityHandle);
}

static void SetEntityRenderColor(int32_t entityHandle, int32_t color) {
	typedef void (*SetEntityRenderColorFn)(int32_t, int32_t);
	static SetEntityRenderColorFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("s2sdk.SetEntityRenderColor", (void**)&__func);
	__func(entityHandle, color);
}

static int8_t GetEntityRenderMode(int32_t entityHandle) {
	typedef int8_t (*GetEntityRenderModeFn)(int32_t);
	static GetEntityRenderModeFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("s2sdk.GetEntityRenderMode", (void**)&__func);
	return __func(entityHandle);
}

static void SetEntityRenderMode(int32_t entityHandle, int8_t renderMode) {
	typedef void (*SetEntityRenderModeFn)(int32_t, int8_t);
	static SetEntityRenderModeFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("s2sdk.SetEntityRenderMode", (void**)&__func);
	__func(entityHandle, renderMode);
}

static int32_t GetEntityHealth(int32_t entityHandle) {
	typedef int32_t (*GetEntityHealthFn)(int32_t);
	static GetEntityHealthFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("s2sdk.GetEntityHealth", (void**)&__func);
	return __func(entityHandle);
}

static void SetEntityHealth(int32_t entityHandle, int32_t health) {
	typedef void (*SetEntityHealthFn)(int32_t, int32_t);
	static SetEntityHealthFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("s2sdk.SetEntityHealth", (void**)&__func);
	__func(entityHandle, health);
}

static int32_t GetTeamEntity(int32_t entityHandle) {
	typedef int32_t (*GetTeamEntityFn)(int32_t);
	static GetTeamEntityFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("s2sdk.GetTeamEntity", (void**)&__func);
	return __func(entityHandle);
}

static void SetTeamEntity(int32_t entityHandle, int32_t team) {
	typedef void (*SetTeamEntityFn)(int32_t, int32_t);
	static SetTeamEntityFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("s2sdk.SetTeamEntity", (void**)&__func);
	__func(entityHandle, team);
}

static int32_t GetEntityOwner(int32_t entityHandle) {
	typedef int32_t (*GetEntityOwnerFn)(int32_t);
	static GetEntityOwnerFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("s2sdk.GetEntityOwner", (void**)&__func);
	return __func(entityHandle);
}

static void SetEntityOwner(int32_t entityHandle, int32_t ownerHandle) {
	typedef void (*SetEntityOwnerFn)(int32_t, int32_t);
	static SetEntityOwnerFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("s2sdk.SetEntityOwner", (void**)&__func);
	__func(entityHandle, ownerHandle);
}

static int32_t GetEntityParent(int32_t entityHandle) {
	typedef int32_t (*GetEntityParentFn)(int32_t);
	static GetEntityParentFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("s2sdk.GetEntityParent", (void**)&__func);
	return __func(entityHandle);
}

static void SetEntityParent(int32_t entityHandle, int32_t parentHandle) {
	typedef void (*SetEntityParentFn)(int32_t, int32_t);
	static SetEntityParentFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("s2sdk.SetEntityParent", (void**)&__func);
	__func(entityHandle, parentHandle);
}

static Vector3 GetEntityAbsOrigin(int32_t entityHandle) {
	typedef Vector3 (*GetEntityAbsOriginFn)(int32_t);
	static GetEntityAbsOriginFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("s2sdk.GetEntityAbsOrigin", (void**)&__func);
	return __func(entityHandle);
}

static void SetEntityAbsOrigin(int32_t entityHandle, Vector3* origin) {
	typedef void (*SetEntityAbsOriginFn)(int32_t, Vector3*);
	static SetEntityAbsOriginFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("s2sdk.SetEntityAbsOrigin", (void**)&__func);
	__func(entityHandle, origin);
}

static Vector3 GetEntityAngRotation(int32_t entityHandle) {
	typedef Vector3 (*GetEntityAngRotationFn)(int32_t);
	static GetEntityAngRotationFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("s2sdk.GetEntityAngRotation", (void**)&__func);
	return __func(entityHandle);
}

static void SetEntityAngRotation(int32_t entityHandle, Vector3* angle) {
	typedef void (*SetEntityAngRotationFn)(int32_t, Vector3*);
	static SetEntityAngRotationFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("s2sdk.SetEntityAngRotation", (void**)&__func);
	__func(entityHandle, angle);
}

static Vector3 GetEntityAbsVelocity(int32_t entityHandle) {
	typedef Vector3 (*GetEntityAbsVelocityFn)(int32_t);
	static GetEntityAbsVelocityFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("s2sdk.GetEntityAbsVelocity", (void**)&__func);
	return __func(entityHandle);
}

static void SetEntityAbsVelocity(int32_t entityHandle, Vector3* velocity) {
	typedef void (*SetEntityAbsVelocityFn)(int32_t, Vector3*);
	static SetEntityAbsVelocityFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("s2sdk.SetEntityAbsVelocity", (void**)&__func);
	__func(entityHandle, velocity);
}

static String GetEntityModel(int32_t entityHandle) {
	typedef String (*GetEntityModelFn)(int32_t);
	static GetEntityModelFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("s2sdk.GetEntityModel", (void**)&__func);
	return __func(entityHandle);
}

static void SetEntityModel(int32_t entityHandle, String* model) {
	typedef void (*SetEntityModelFn)(int32_t, String*);
	static SetEntityModelFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("s2sdk.SetEntityModel", (void**)&__func);
	__func(entityHandle, model);
}

static float GetEntityWaterLevel(int32_t entityHandle) {
	typedef float (*GetEntityWaterLevelFn)(int32_t);
	static GetEntityWaterLevelFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("s2sdk.GetEntityWaterLevel", (void**)&__func);
	return __func(entityHandle);
}

static int32_t GetEntityGroundEntity(int32_t entityHandle) {
	typedef int32_t (*GetEntityGroundEntityFn)(int32_t);
	static GetEntityGroundEntityFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("s2sdk.GetEntityGroundEntity", (void**)&__func);
	return __func(entityHandle);
}

static int32_t GetEntityEffects(int32_t entityHandle) {
	typedef int32_t (*GetEntityEffectsFn)(int32_t);
	static GetEntityEffectsFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("s2sdk.GetEntityEffects", (void**)&__func);
	return __func(entityHandle);
}

static void TeleportEntity(int32_t entityHandle, uintptr_t origin, uintptr_t angles, uintptr_t velocity) {
	typedef void (*TeleportEntityFn)(int32_t, uintptr_t, uintptr_t, uintptr_t);
	static TeleportEntityFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("s2sdk.TeleportEntity", (void**)&__func);
	__func(entityHandle, origin, angles, velocity);
}

static void AcceptInput(int32_t entityHandle, String* inputName, int32_t activatorHandle, int32_t callerHandle, Variant* value, int32_t type, int32_t outputId) {
	typedef void (*AcceptInputFn)(int32_t, String*, int32_t, int32_t, Variant*, int32_t, int32_t);
	static AcceptInputFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("s2sdk.AcceptInput", (void**)&__func);
	__func(entityHandle, inputName, activatorHandle, callerHandle, value, type, outputId);
}

static int32_t HookEvent(String* name, void* pCallback, uint8_t type) {
	typedef int32_t (*HookEventFn)(String*, void*, uint8_t);
	static HookEventFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("s2sdk.HookEvent", (void**)&__func);
	return __func(name, pCallback, type);
}

static int32_t UnhookEvent(String* name, void* pCallback, uint8_t type) {
	typedef int32_t (*UnhookEventFn)(String*, void*, uint8_t);
	static UnhookEventFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("s2sdk.UnhookEvent", (void**)&__func);
	return __func(name, pCallback, type);
}

static uintptr_t CreateEvent(String* name, bool force) {
	typedef uintptr_t (*CreateEventFn)(String*, bool);
	static CreateEventFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("s2sdk.CreateEvent", (void**)&__func);
	return __func(name, force);
}

static void FireEvent(uintptr_t pInfo, bool bDontBroadcast) {
	typedef void (*FireEventFn)(uintptr_t, bool);
	static FireEventFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("s2sdk.FireEvent", (void**)&__func);
	__func(pInfo, bDontBroadcast);
}

static void FireEventToClient(uintptr_t pInfo, int32_t playerSlot) {
	typedef void (*FireEventToClientFn)(uintptr_t, int32_t);
	static FireEventToClientFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("s2sdk.FireEventToClient", (void**)&__func);
	__func(pInfo, playerSlot);
}

static void CancelCreatedEvent(uintptr_t pInfo) {
	typedef void (*CancelCreatedEventFn)(uintptr_t);
	static CancelCreatedEventFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("s2sdk.CancelCreatedEvent", (void**)&__func);
	__func(pInfo);
}

static bool GetEventBool(uintptr_t pInfo, String* key) {
	typedef bool (*GetEventBoolFn)(uintptr_t, String*);
	static GetEventBoolFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("s2sdk.GetEventBool", (void**)&__func);
	return __func(pInfo, key);
}

static float GetEventFloat(uintptr_t pInfo, String* key) {
	typedef float (*GetEventFloatFn)(uintptr_t, String*);
	static GetEventFloatFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("s2sdk.GetEventFloat", (void**)&__func);
	return __func(pInfo, key);
}

static int32_t GetEventInt(uintptr_t pInfo, String* key) {
	typedef int32_t (*GetEventIntFn)(uintptr_t, String*);
	static GetEventIntFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("s2sdk.GetEventInt", (void**)&__func);
	return __func(pInfo, key);
}

static uint64_t GetEventUInt64(uintptr_t pInfo, String* key) {
	typedef uint64_t (*GetEventUInt64Fn)(uintptr_t, String*);
	static GetEventUInt64Fn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("s2sdk.GetEventUInt64", (void**)&__func);
	return __func(pInfo, key);
}

static String GetEventString(uintptr_t pInfo, String* key) {
	typedef String (*GetEventStringFn)(uintptr_t, String*);
	static GetEventStringFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("s2sdk.GetEventString", (void**)&__func);
	return __func(pInfo, key);
}

static uintptr_t GetEventPtr(uintptr_t pInfo, String* key) {
	typedef uintptr_t (*GetEventPtrFn)(uintptr_t, String*);
	static GetEventPtrFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("s2sdk.GetEventPtr", (void**)&__func);
	return __func(pInfo, key);
}

static uintptr_t GetEventPlayerController(uintptr_t pInfo, String* key) {
	typedef uintptr_t (*GetEventPlayerControllerFn)(uintptr_t, String*);
	static GetEventPlayerControllerFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("s2sdk.GetEventPlayerController", (void**)&__func);
	return __func(pInfo, key);
}

static int32_t GetEventPlayerIndex(uintptr_t pInfo, String* key) {
	typedef int32_t (*GetEventPlayerIndexFn)(uintptr_t, String*);
	static GetEventPlayerIndexFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("s2sdk.GetEventPlayerIndex", (void**)&__func);
	return __func(pInfo, key);
}

static uintptr_t GetEventPlayerPawn(uintptr_t pInfo, String* key) {
	typedef uintptr_t (*GetEventPlayerPawnFn)(uintptr_t, String*);
	static GetEventPlayerPawnFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("s2sdk.GetEventPlayerPawn", (void**)&__func);
	return __func(pInfo, key);
}

static uintptr_t GetEventEntity(uintptr_t pInfo, String* key) {
	typedef uintptr_t (*GetEventEntityFn)(uintptr_t, String*);
	static GetEventEntityFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("s2sdk.GetEventEntity", (void**)&__func);
	return __func(pInfo, key);
}

static int32_t GetEventEntityIndex(uintptr_t pInfo, String* key) {
	typedef int32_t (*GetEventEntityIndexFn)(uintptr_t, String*);
	static GetEventEntityIndexFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("s2sdk.GetEventEntityIndex", (void**)&__func);
	return __func(pInfo, key);
}

static int32_t GetEventEntityHandle(uintptr_t pInfo, String* key) {
	typedef int32_t (*GetEventEntityHandleFn)(uintptr_t, String*);
	static GetEventEntityHandleFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("s2sdk.GetEventEntityHandle", (void**)&__func);
	return __func(pInfo, key);
}

static String GetEventName(uintptr_t pInfo) {
	typedef String (*GetEventNameFn)(uintptr_t);
	static GetEventNameFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("s2sdk.GetEventName", (void**)&__func);
	return __func(pInfo);
}

static void SetEventBool(uintptr_t pInfo, String* key, bool value) {
	typedef void (*SetEventBoolFn)(uintptr_t, String*, bool);
	static SetEventBoolFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("s2sdk.SetEventBool", (void**)&__func);
	__func(pInfo, key, value);
}

static void SetEventFloat(uintptr_t pInfo, String* key, float value) {
	typedef void (*SetEventFloatFn)(uintptr_t, String*, float);
	static SetEventFloatFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("s2sdk.SetEventFloat", (void**)&__func);
	__func(pInfo, key, value);
}

static void SetEventInt(uintptr_t pInfo, String* key, int32_t value) {
	typedef void (*SetEventIntFn)(uintptr_t, String*, int32_t);
	static SetEventIntFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("s2sdk.SetEventInt", (void**)&__func);
	__func(pInfo, key, value);
}

static void SetEventUInt64(uintptr_t pInfo, String* key, uint64_t value) {
	typedef void (*SetEventUInt64Fn)(uintptr_t, String*, uint64_t);
	static SetEventUInt64Fn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("s2sdk.SetEventUInt64", (void**)&__func);
	__func(pInfo, key, value);
}

static void SetEventString(uintptr_t pInfo, String* key, String* value) {
	typedef void (*SetEventStringFn)(uintptr_t, String*, String*);
	static SetEventStringFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("s2sdk.SetEventString", (void**)&__func);
	__func(pInfo, key, value);
}

static void SetEventPtr(uintptr_t pInfo, String* key, uintptr_t value) {
	typedef void (*SetEventPtrFn)(uintptr_t, String*, uintptr_t);
	static SetEventPtrFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("s2sdk.SetEventPtr", (void**)&__func);
	__func(pInfo, key, value);
}

static void SetEventPlayerController(uintptr_t pInfo, String* key, uintptr_t value) {
	typedef void (*SetEventPlayerControllerFn)(uintptr_t, String*, uintptr_t);
	static SetEventPlayerControllerFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("s2sdk.SetEventPlayerController", (void**)&__func);
	__func(pInfo, key, value);
}

static void SetEventPlayerIndex(uintptr_t pInfo, String* key, int32_t value) {
	typedef void (*SetEventPlayerIndexFn)(uintptr_t, String*, int32_t);
	static SetEventPlayerIndexFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("s2sdk.SetEventPlayerIndex", (void**)&__func);
	__func(pInfo, key, value);
}

static void SetEventEntity(uintptr_t pInfo, String* key, uintptr_t value) {
	typedef void (*SetEventEntityFn)(uintptr_t, String*, uintptr_t);
	static SetEventEntityFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("s2sdk.SetEventEntity", (void**)&__func);
	__func(pInfo, key, value);
}

static void SetEventEntityIndex(uintptr_t pInfo, String* key, int32_t value) {
	typedef void (*SetEventEntityIndexFn)(uintptr_t, String*, int32_t);
	static SetEventEntityIndexFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("s2sdk.SetEventEntityIndex", (void**)&__func);
	__func(pInfo, key, value);
}

static void SetEventEntityHandle(uintptr_t pInfo, String* key, int32_t value) {
	typedef void (*SetEventEntityHandleFn)(uintptr_t, String*, int32_t);
	static SetEventEntityHandleFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("s2sdk.SetEventEntityHandle", (void**)&__func);
	__func(pInfo, key, value);
}

static void SetEventBroadcast(uintptr_t pInfo, bool dontBroadcast) {
	typedef void (*SetEventBroadcastFn)(uintptr_t, bool);
	static SetEventBroadcastFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("s2sdk.SetEventBroadcast", (void**)&__func);
	__func(pInfo, dontBroadcast);
}

static int32_t LoadEventsFromFile(String* path, bool searchAll) {
	typedef int32_t (*LoadEventsFromFileFn)(String*, bool);
	static LoadEventsFromFileFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("s2sdk.LoadEventsFromFile", (void**)&__func);
	return __func(path, searchAll);
}

static void CloseGameConfigFile(uintptr_t pGameConfig) {
	typedef void (*CloseGameConfigFileFn)(uintptr_t);
	static CloseGameConfigFileFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("s2sdk.CloseGameConfigFile", (void**)&__func);
	__func(pGameConfig);
}

static uintptr_t LoadGameConfigFile(String* file) {
	typedef uintptr_t (*LoadGameConfigFileFn)(String*);
	static LoadGameConfigFileFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("s2sdk.LoadGameConfigFile", (void**)&__func);
	return __func(file);
}

static String GetGameConfigPath(uintptr_t pGameConfig) {
	typedef String (*GetGameConfigPathFn)(uintptr_t);
	static GetGameConfigPathFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("s2sdk.GetGameConfigPath", (void**)&__func);
	return __func(pGameConfig);
}

static String GetGameConfigLibrary(uintptr_t pGameConfig, String* name) {
	typedef String (*GetGameConfigLibraryFn)(uintptr_t, String*);
	static GetGameConfigLibraryFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("s2sdk.GetGameConfigLibrary", (void**)&__func);
	return __func(pGameConfig, name);
}

static String GetGameConfigSignature(uintptr_t pGameConfig, String* name) {
	typedef String (*GetGameConfigSignatureFn)(uintptr_t, String*);
	static GetGameConfigSignatureFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("s2sdk.GetGameConfigSignature", (void**)&__func);
	return __func(pGameConfig, name);
}

static String GetGameConfigSymbol(uintptr_t pGameConfig, String* name) {
	typedef String (*GetGameConfigSymbolFn)(uintptr_t, String*);
	static GetGameConfigSymbolFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("s2sdk.GetGameConfigSymbol", (void**)&__func);
	return __func(pGameConfig, name);
}

static String GetGameConfigPatch(uintptr_t pGameConfig, String* name) {
	typedef String (*GetGameConfigPatchFn)(uintptr_t, String*);
	static GetGameConfigPatchFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("s2sdk.GetGameConfigPatch", (void**)&__func);
	return __func(pGameConfig, name);
}

static int32_t GetGameConfigOffset(uintptr_t pGameConfig, String* name) {
	typedef int32_t (*GetGameConfigOffsetFn)(uintptr_t, String*);
	static GetGameConfigOffsetFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("s2sdk.GetGameConfigOffset", (void**)&__func);
	return __func(pGameConfig, name);
}

static uintptr_t GetGameConfigAddress(uintptr_t pGameConfig, String* name) {
	typedef uintptr_t (*GetGameConfigAddressFn)(uintptr_t, String*);
	static GetGameConfigAddressFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("s2sdk.GetGameConfigAddress", (void**)&__func);
	return __func(pGameConfig, name);
}

static uintptr_t GetGameConfigMemSig(uintptr_t pGameConfig, String* name) {
	typedef uintptr_t (*GetGameConfigMemSigFn)(uintptr_t, String*);
	static GetGameConfigMemSigFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("s2sdk.GetGameConfigMemSig", (void**)&__func);
	return __func(pGameConfig, name);
}

static int32_t RegisterLoggingChannel(String* name, int32_t iFlags, int32_t verbosity, int32_t color) {
	typedef int32_t (*RegisterLoggingChannelFn)(String*, int32_t, int32_t, int32_t);
	static RegisterLoggingChannelFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("s2sdk.RegisterLoggingChannel", (void**)&__func);
	return __func(name, iFlags, verbosity, color);
}

static void AddLoggerTagToChannel(int32_t channelID, String* tagName) {
	typedef void (*AddLoggerTagToChannelFn)(int32_t, String*);
	static AddLoggerTagToChannelFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("s2sdk.AddLoggerTagToChannel", (void**)&__func);
	__func(channelID, tagName);
}

static bool HasLoggerTag(int32_t channelID, String* tag) {
	typedef bool (*HasLoggerTagFn)(int32_t, String*);
	static HasLoggerTagFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("s2sdk.HasLoggerTag", (void**)&__func);
	return __func(channelID, tag);
}

static bool IsLoggerChannelEnabledBySeverity(int32_t channelID, int32_t severity) {
	typedef bool (*IsLoggerChannelEnabledBySeverityFn)(int32_t, int32_t);
	static IsLoggerChannelEnabledBySeverityFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("s2sdk.IsLoggerChannelEnabledBySeverity", (void**)&__func);
	return __func(channelID, severity);
}

static bool IsLoggerChannelEnabledByVerbosity(int32_t channelID, int32_t verbosity) {
	typedef bool (*IsLoggerChannelEnabledByVerbosityFn)(int32_t, int32_t);
	static IsLoggerChannelEnabledByVerbosityFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("s2sdk.IsLoggerChannelEnabledByVerbosity", (void**)&__func);
	return __func(channelID, verbosity);
}

static int32_t GetLoggerChannelVerbosity(int32_t channelID) {
	typedef int32_t (*GetLoggerChannelVerbosityFn)(int32_t);
	static GetLoggerChannelVerbosityFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("s2sdk.GetLoggerChannelVerbosity", (void**)&__func);
	return __func(channelID);
}

static void SetLoggerChannelVerbosity(int32_t channelID, int32_t verbosity) {
	typedef void (*SetLoggerChannelVerbosityFn)(int32_t, int32_t);
	static SetLoggerChannelVerbosityFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("s2sdk.SetLoggerChannelVerbosity", (void**)&__func);
	__func(channelID, verbosity);
}

static void SetLoggerChannelVerbosityByName(int32_t channelID, String* name, int32_t verbosity) {
	typedef void (*SetLoggerChannelVerbosityByNameFn)(int32_t, String*, int32_t);
	static SetLoggerChannelVerbosityByNameFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("s2sdk.SetLoggerChannelVerbosityByName", (void**)&__func);
	__func(channelID, name, verbosity);
}

static void SetLoggerChannelVerbosityByTag(int32_t channelID, String* tag, int32_t verbosity) {
	typedef void (*SetLoggerChannelVerbosityByTagFn)(int32_t, String*, int32_t);
	static SetLoggerChannelVerbosityByTagFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("s2sdk.SetLoggerChannelVerbosityByTag", (void**)&__func);
	__func(channelID, tag, verbosity);
}

static int32_t GetLoggerChannelColor(int32_t channelID) {
	typedef int32_t (*GetLoggerChannelColorFn)(int32_t);
	static GetLoggerChannelColorFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("s2sdk.GetLoggerChannelColor", (void**)&__func);
	return __func(channelID);
}

static void SetLoggerChannelColor(int32_t channelID, int32_t color) {
	typedef void (*SetLoggerChannelColorFn)(int32_t, int32_t);
	static SetLoggerChannelColorFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("s2sdk.SetLoggerChannelColor", (void**)&__func);
	__func(channelID, color);
}

static int32_t GetLoggerChannelFlags(int32_t channelID) {
	typedef int32_t (*GetLoggerChannelFlagsFn)(int32_t);
	static GetLoggerChannelFlagsFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("s2sdk.GetLoggerChannelFlags", (void**)&__func);
	return __func(channelID);
}

static void SetLoggerChannelFlags(int32_t channelID, int32_t eFlags) {
	typedef void (*SetLoggerChannelFlagsFn)(int32_t, int32_t);
	static SetLoggerChannelFlagsFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("s2sdk.SetLoggerChannelFlags", (void**)&__func);
	__func(channelID, eFlags);
}

static int32_t Log(int32_t channelID, int32_t severity, String* message) {
	typedef int32_t (*LogFn)(int32_t, int32_t, String*);
	static LogFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("s2sdk.Log", (void**)&__func);
	return __func(channelID, severity, message);
}

static int32_t LogColored(int32_t channelID, int32_t severity, int32_t color, String* message) {
	typedef int32_t (*LogColoredFn)(int32_t, int32_t, int32_t, String*);
	static LogColoredFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("s2sdk.LogColored", (void**)&__func);
	return __func(channelID, severity, color, message);
}

static int32_t LogFull(int32_t channelID, int32_t severity, String* file, int32_t line, String* function, String* message) {
	typedef int32_t (*LogFullFn)(int32_t, int32_t, String*, int32_t, String*, String*);
	static LogFullFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("s2sdk.LogFull", (void**)&__func);
	return __func(channelID, severity, file, line, function, message);
}

static int32_t LogFullColored(int32_t channelID, int32_t severity, String* file, int32_t line, String* function, int32_t color, String* message) {
	typedef int32_t (*LogFullColoredFn)(int32_t, int32_t, String*, int32_t, String*, int32_t, String*);
	static LogFullColoredFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("s2sdk.LogFullColored", (void**)&__func);
	return __func(channelID, severity, file, line, function, color, message);
}

static int32_t GetSchemaOffset(String* className, String* memberName) {
	typedef int32_t (*GetSchemaOffsetFn)(String*, String*);
	static GetSchemaOffsetFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("s2sdk.GetSchemaOffset", (void**)&__func);
	return __func(className, memberName);
}

static int32_t GetSchemaChainOffset(String* className) {
	typedef int32_t (*GetSchemaChainOffsetFn)(String*);
	static GetSchemaChainOffsetFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("s2sdk.GetSchemaChainOffset", (void**)&__func);
	return __func(className);
}

static bool IsSchemaFieldNetworked(String* className, String* memberName) {
	typedef bool (*IsSchemaFieldNetworkedFn)(String*, String*);
	static IsSchemaFieldNetworkedFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("s2sdk.IsSchemaFieldNetworked", (void**)&__func);
	return __func(className, memberName);
}

static int32_t GetSchemaClassSize(String* className) {
	typedef int32_t (*GetSchemaClassSizeFn)(String*);
	static GetSchemaClassSizeFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("s2sdk.GetSchemaClassSize", (void**)&__func);
	return __func(className);
}

static int64_t GetEntData2(uintptr_t entity, int32_t offset, int32_t size) {
	typedef int64_t (*GetEntData2Fn)(uintptr_t, int32_t, int32_t);
	static GetEntData2Fn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("s2sdk.GetEntData2", (void**)&__func);
	return __func(entity, offset, size);
}

static void SetEntData2(uintptr_t entity, int32_t offset, int64_t value, int32_t size, bool changeState, int32_t chainOffset) {
	typedef void (*SetEntData2Fn)(uintptr_t, int32_t, int64_t, int32_t, bool, int32_t);
	static SetEntData2Fn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("s2sdk.SetEntData2", (void**)&__func);
	__func(entity, offset, value, size, changeState, chainOffset);
}

static double GetEntDataFloat2(uintptr_t entity, int32_t offset, int32_t size) {
	typedef double (*GetEntDataFloat2Fn)(uintptr_t, int32_t, int32_t);
	static GetEntDataFloat2Fn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("s2sdk.GetEntDataFloat2", (void**)&__func);
	return __func(entity, offset, size);
}

static void SetEntDataFloat2(uintptr_t entity, int32_t offset, double value, int32_t size, bool changeState, int32_t chainOffset) {
	typedef void (*SetEntDataFloat2Fn)(uintptr_t, int32_t, double, int32_t, bool, int32_t);
	static SetEntDataFloat2Fn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("s2sdk.SetEntDataFloat2", (void**)&__func);
	__func(entity, offset, value, size, changeState, chainOffset);
}

static String GetEntDataString2(uintptr_t entity, int32_t offset) {
	typedef String (*GetEntDataString2Fn)(uintptr_t, int32_t);
	static GetEntDataString2Fn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("s2sdk.GetEntDataString2", (void**)&__func);
	return __func(entity, offset);
}

static void SetEntDataString2(uintptr_t entity, int32_t offset, String* value, bool changeState, int32_t chainOffset) {
	typedef void (*SetEntDataString2Fn)(uintptr_t, int32_t, String*, bool, int32_t);
	static SetEntDataString2Fn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("s2sdk.SetEntDataString2", (void**)&__func);
	__func(entity, offset, value, changeState, chainOffset);
}

static Vector3 GetEntDataVector2(uintptr_t entity, int32_t offset) {
	typedef Vector3 (*GetEntDataVector2Fn)(uintptr_t, int32_t);
	static GetEntDataVector2Fn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("s2sdk.GetEntDataVector2", (void**)&__func);
	return __func(entity, offset);
}

static void SetEntDataVector2(uintptr_t entity, int32_t offset, Vector3* value, bool changeState, int32_t chainOffset) {
	typedef void (*SetEntDataVector2Fn)(uintptr_t, int32_t, Vector3*, bool, int32_t);
	static SetEntDataVector2Fn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("s2sdk.SetEntDataVector2", (void**)&__func);
	__func(entity, offset, value, changeState, chainOffset);
}

static int32_t GetEntDataEnt2(uintptr_t entity, int32_t offset) {
	typedef int32_t (*GetEntDataEnt2Fn)(uintptr_t, int32_t);
	static GetEntDataEnt2Fn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("s2sdk.GetEntDataEnt2", (void**)&__func);
	return __func(entity, offset);
}

static void SetEntDataEnt2(uintptr_t entity, int32_t offset, int32_t value, bool changeState, int32_t chainOffset) {
	typedef void (*SetEntDataEnt2Fn)(uintptr_t, int32_t, int32_t, bool, int32_t);
	static SetEntDataEnt2Fn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("s2sdk.SetEntDataEnt2", (void**)&__func);
	__func(entity, offset, value, changeState, chainOffset);
}

static void ChangeEntityState2(uintptr_t entity, int32_t offset, int32_t chainOffset) {
	typedef void (*ChangeEntityState2Fn)(uintptr_t, int32_t, int32_t);
	static ChangeEntityState2Fn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("s2sdk.ChangeEntityState2", (void**)&__func);
	__func(entity, offset, chainOffset);
}

static int64_t GetEntData(int32_t entityHandle, int32_t offset, int32_t size) {
	typedef int64_t (*GetEntDataFn)(int32_t, int32_t, int32_t);
	static GetEntDataFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("s2sdk.GetEntData", (void**)&__func);
	return __func(entityHandle, offset, size);
}

static void SetEntData(int32_t entityHandle, int32_t offset, int64_t value, int32_t size, bool changeState, int32_t chainOffset) {
	typedef void (*SetEntDataFn)(int32_t, int32_t, int64_t, int32_t, bool, int32_t);
	static SetEntDataFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("s2sdk.SetEntData", (void**)&__func);
	__func(entityHandle, offset, value, size, changeState, chainOffset);
}

static double GetEntDataFloat(int32_t entityHandle, int32_t offset, int32_t size) {
	typedef double (*GetEntDataFloatFn)(int32_t, int32_t, int32_t);
	static GetEntDataFloatFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("s2sdk.GetEntDataFloat", (void**)&__func);
	return __func(entityHandle, offset, size);
}

static void SetEntDataFloat(int32_t entityHandle, int32_t offset, double value, int32_t size, bool changeState, int32_t chainOffset) {
	typedef void (*SetEntDataFloatFn)(int32_t, int32_t, double, int32_t, bool, int32_t);
	static SetEntDataFloatFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("s2sdk.SetEntDataFloat", (void**)&__func);
	__func(entityHandle, offset, value, size, changeState, chainOffset);
}

static String GetEntDataString(int32_t entityHandle, int32_t offset) {
	typedef String (*GetEntDataStringFn)(int32_t, int32_t);
	static GetEntDataStringFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("s2sdk.GetEntDataString", (void**)&__func);
	return __func(entityHandle, offset);
}

static void SetEntDataString(int32_t entityHandle, int32_t offset, String* value, bool changeState, int32_t chainOffset) {
	typedef void (*SetEntDataStringFn)(int32_t, int32_t, String*, bool, int32_t);
	static SetEntDataStringFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("s2sdk.SetEntDataString", (void**)&__func);
	__func(entityHandle, offset, value, changeState, chainOffset);
}

static Vector3 GetEntDataVector(int32_t entityHandle, int32_t offset) {
	typedef Vector3 (*GetEntDataVectorFn)(int32_t, int32_t);
	static GetEntDataVectorFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("s2sdk.GetEntDataVector", (void**)&__func);
	return __func(entityHandle, offset);
}

static void SetEntDataVector(int32_t entityHandle, int32_t offset, Vector3* value, bool changeState, int32_t chainOffset) {
	typedef void (*SetEntDataVectorFn)(int32_t, int32_t, Vector3*, bool, int32_t);
	static SetEntDataVectorFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("s2sdk.SetEntDataVector", (void**)&__func);
	__func(entityHandle, offset, value, changeState, chainOffset);
}

static int32_t GetEntDataEnt(int32_t entityHandle, int32_t offset) {
	typedef int32_t (*GetEntDataEntFn)(int32_t, int32_t);
	static GetEntDataEntFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("s2sdk.GetEntDataEnt", (void**)&__func);
	return __func(entityHandle, offset);
}

static void SetEntDataEnt(int32_t entityHandle, int32_t offset, int32_t value, bool changeState, int32_t chainOffset) {
	typedef void (*SetEntDataEntFn)(int32_t, int32_t, int32_t, bool, int32_t);
	static SetEntDataEntFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("s2sdk.SetEntDataEnt", (void**)&__func);
	__func(entityHandle, offset, value, changeState, chainOffset);
}

static void ChangeEntityState(int32_t entityHandle, int32_t offset, int32_t chainOffset) {
	typedef void (*ChangeEntityStateFn)(int32_t, int32_t, int32_t);
	static ChangeEntityStateFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("s2sdk.ChangeEntityState", (void**)&__func);
	__func(entityHandle, offset, chainOffset);
}

static int32_t GetEntSchemaArraySize2(uintptr_t entity, String* className, String* memberName) {
	typedef int32_t (*GetEntSchemaArraySize2Fn)(uintptr_t, String*, String*);
	static GetEntSchemaArraySize2Fn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("s2sdk.GetEntSchemaArraySize2", (void**)&__func);
	return __func(entity, className, memberName);
}

static int64_t GetEntSchema2(uintptr_t entity, String* className, String* memberName, int32_t element) {
	typedef int64_t (*GetEntSchema2Fn)(uintptr_t, String*, String*, int32_t);
	static GetEntSchema2Fn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("s2sdk.GetEntSchema2", (void**)&__func);
	return __func(entity, className, memberName, element);
}

static void SetEntSchema2(uintptr_t entity, String* className, String* memberName, int64_t value, bool changeState, int32_t element) {
	typedef void (*SetEntSchema2Fn)(uintptr_t, String*, String*, int64_t, bool, int32_t);
	static SetEntSchema2Fn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("s2sdk.SetEntSchema2", (void**)&__func);
	__func(entity, className, memberName, value, changeState, element);
}

static double GetEntSchemaFloat2(uintptr_t entity, String* className, String* memberName, int32_t element) {
	typedef double (*GetEntSchemaFloat2Fn)(uintptr_t, String*, String*, int32_t);
	static GetEntSchemaFloat2Fn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("s2sdk.GetEntSchemaFloat2", (void**)&__func);
	return __func(entity, className, memberName, element);
}

static void SetEntSchemaFloat2(uintptr_t entity, String* className, String* memberName, double value, bool changeState, int32_t element) {
	typedef void (*SetEntSchemaFloat2Fn)(uintptr_t, String*, String*, double, bool, int32_t);
	static SetEntSchemaFloat2Fn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("s2sdk.SetEntSchemaFloat2", (void**)&__func);
	__func(entity, className, memberName, value, changeState, element);
}

static String GetEntSchemaString2(uintptr_t entity, String* className, String* memberName, int32_t element) {
	typedef String (*GetEntSchemaString2Fn)(uintptr_t, String*, String*, int32_t);
	static GetEntSchemaString2Fn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("s2sdk.GetEntSchemaString2", (void**)&__func);
	return __func(entity, className, memberName, element);
}

static void SetEntSchemaString2(uintptr_t entity, String* className, String* memberName, String* value, bool changeState, int32_t element) {
	typedef void (*SetEntSchemaString2Fn)(uintptr_t, String*, String*, String*, bool, int32_t);
	static SetEntSchemaString2Fn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("s2sdk.SetEntSchemaString2", (void**)&__func);
	__func(entity, className, memberName, value, changeState, element);
}

static Vector3 GetEntSchemaVector3D2(uintptr_t entity, String* className, String* memberName, int32_t element) {
	typedef Vector3 (*GetEntSchemaVector3D2Fn)(uintptr_t, String*, String*, int32_t);
	static GetEntSchemaVector3D2Fn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("s2sdk.GetEntSchemaVector3D2", (void**)&__func);
	return __func(entity, className, memberName, element);
}

static void SetEntSchemaVector3D2(uintptr_t entity, String* className, String* memberName, Vector3* value, bool changeState, int32_t element) {
	typedef void (*SetEntSchemaVector3D2Fn)(uintptr_t, String*, String*, Vector3*, bool, int32_t);
	static SetEntSchemaVector3D2Fn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("s2sdk.SetEntSchemaVector3D2", (void**)&__func);
	__func(entity, className, memberName, value, changeState, element);
}

static Vector2 GetEntSchemaVector2D2(uintptr_t entity, String* className, String* memberName, int32_t element) {
	typedef Vector2 (*GetEntSchemaVector2D2Fn)(uintptr_t, String*, String*, int32_t);
	static GetEntSchemaVector2D2Fn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("s2sdk.GetEntSchemaVector2D2", (void**)&__func);
	return __func(entity, className, memberName, element);
}

static void SetEntSchemaVector2D2(uintptr_t entity, String* className, String* memberName, Vector2* value, bool changeState, int32_t element) {
	typedef void (*SetEntSchemaVector2D2Fn)(uintptr_t, String*, String*, Vector2*, bool, int32_t);
	static SetEntSchemaVector2D2Fn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("s2sdk.SetEntSchemaVector2D2", (void**)&__func);
	__func(entity, className, memberName, value, changeState, element);
}

static Vector4 GetEntSchemaVector4D2(uintptr_t entity, String* className, String* memberName, int32_t element) {
	typedef Vector4 (*GetEntSchemaVector4D2Fn)(uintptr_t, String*, String*, int32_t);
	static GetEntSchemaVector4D2Fn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("s2sdk.GetEntSchemaVector4D2", (void**)&__func);
	return __func(entity, className, memberName, element);
}

static void SetEntSchemaVector4D2(uintptr_t entity, String* className, String* memberName, Vector4* value, bool changeState, int32_t element) {
	typedef void (*SetEntSchemaVector4D2Fn)(uintptr_t, String*, String*, Vector4*, bool, int32_t);
	static SetEntSchemaVector4D2Fn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("s2sdk.SetEntSchemaVector4D2", (void**)&__func);
	__func(entity, className, memberName, value, changeState, element);
}

static int32_t GetEntSchemaEnt2(uintptr_t entity, String* className, String* memberName, int32_t element) {
	typedef int32_t (*GetEntSchemaEnt2Fn)(uintptr_t, String*, String*, int32_t);
	static GetEntSchemaEnt2Fn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("s2sdk.GetEntSchemaEnt2", (void**)&__func);
	return __func(entity, className, memberName, element);
}

static void SetEntSchemaEnt2(uintptr_t entity, String* className, String* memberName, int32_t value, bool changeState, int32_t element) {
	typedef void (*SetEntSchemaEnt2Fn)(uintptr_t, String*, String*, int32_t, bool, int32_t);
	static SetEntSchemaEnt2Fn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("s2sdk.SetEntSchemaEnt2", (void**)&__func);
	__func(entity, className, memberName, value, changeState, element);
}

static void NetworkStateChanged2(uintptr_t entity, String* className, String* memberName) {
	typedef void (*NetworkStateChanged2Fn)(uintptr_t, String*, String*);
	static NetworkStateChanged2Fn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("s2sdk.NetworkStateChanged2", (void**)&__func);
	__func(entity, className, memberName);
}

static int32_t GetEntSchemaArraySize(int32_t entityHandle, String* className, String* memberName) {
	typedef int32_t (*GetEntSchemaArraySizeFn)(int32_t, String*, String*);
	static GetEntSchemaArraySizeFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("s2sdk.GetEntSchemaArraySize", (void**)&__func);
	return __func(entityHandle, className, memberName);
}

static int64_t GetEntSchema(int32_t entityHandle, String* className, String* memberName, int32_t element) {
	typedef int64_t (*GetEntSchemaFn)(int32_t, String*, String*, int32_t);
	static GetEntSchemaFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("s2sdk.GetEntSchema", (void**)&__func);
	return __func(entityHandle, className, memberName, element);
}

static void SetEntSchema(int32_t entityHandle, String* className, String* memberName, int64_t value, bool changeState, int32_t element) {
	typedef void (*SetEntSchemaFn)(int32_t, String*, String*, int64_t, bool, int32_t);
	static SetEntSchemaFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("s2sdk.SetEntSchema", (void**)&__func);
	__func(entityHandle, className, memberName, value, changeState, element);
}

static double GetEntSchemaFloat(int32_t entityHandle, String* className, String* memberName, int32_t element) {
	typedef double (*GetEntSchemaFloatFn)(int32_t, String*, String*, int32_t);
	static GetEntSchemaFloatFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("s2sdk.GetEntSchemaFloat", (void**)&__func);
	return __func(entityHandle, className, memberName, element);
}

static void SetEntSchemaFloat(int32_t entityHandle, String* className, String* memberName, double value, bool changeState, int32_t element) {
	typedef void (*SetEntSchemaFloatFn)(int32_t, String*, String*, double, bool, int32_t);
	static SetEntSchemaFloatFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("s2sdk.SetEntSchemaFloat", (void**)&__func);
	__func(entityHandle, className, memberName, value, changeState, element);
}

static String GetEntSchemaString(int32_t entityHandle, String* className, String* memberName, int32_t element) {
	typedef String (*GetEntSchemaStringFn)(int32_t, String*, String*, int32_t);
	static GetEntSchemaStringFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("s2sdk.GetEntSchemaString", (void**)&__func);
	return __func(entityHandle, className, memberName, element);
}

static void SetEntSchemaString(int32_t entityHandle, String* className, String* memberName, String* value, bool changeState, int32_t element) {
	typedef void (*SetEntSchemaStringFn)(int32_t, String*, String*, String*, bool, int32_t);
	static SetEntSchemaStringFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("s2sdk.SetEntSchemaString", (void**)&__func);
	__func(entityHandle, className, memberName, value, changeState, element);
}

static Vector3 GetEntSchemaVector3D(int32_t entityHandle, String* className, String* memberName, int32_t element) {
	typedef Vector3 (*GetEntSchemaVector3DFn)(int32_t, String*, String*, int32_t);
	static GetEntSchemaVector3DFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("s2sdk.GetEntSchemaVector3D", (void**)&__func);
	return __func(entityHandle, className, memberName, element);
}

static void SetEntSchemaVector3D(int32_t entityHandle, String* className, String* memberName, Vector3* value, bool changeState, int32_t element) {
	typedef void (*SetEntSchemaVector3DFn)(int32_t, String*, String*, Vector3*, bool, int32_t);
	static SetEntSchemaVector3DFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("s2sdk.SetEntSchemaVector3D", (void**)&__func);
	__func(entityHandle, className, memberName, value, changeState, element);
}

static Vector2 GetEntSchemaVector2D(int32_t entityHandle, String* className, String* memberName, int32_t element) {
	typedef Vector2 (*GetEntSchemaVector2DFn)(int32_t, String*, String*, int32_t);
	static GetEntSchemaVector2DFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("s2sdk.GetEntSchemaVector2D", (void**)&__func);
	return __func(entityHandle, className, memberName, element);
}

static void SetEntSchemaVector2D(int32_t entityHandle, String* className, String* memberName, Vector2* value, bool changeState, int32_t element) {
	typedef void (*SetEntSchemaVector2DFn)(int32_t, String*, String*, Vector2*, bool, int32_t);
	static SetEntSchemaVector2DFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("s2sdk.SetEntSchemaVector2D", (void**)&__func);
	__func(entityHandle, className, memberName, value, changeState, element);
}

static Vector4 GetEntSchemaVector4D(int32_t entityHandle, String* className, String* memberName, int32_t element) {
	typedef Vector4 (*GetEntSchemaVector4DFn)(int32_t, String*, String*, int32_t);
	static GetEntSchemaVector4DFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("s2sdk.GetEntSchemaVector4D", (void**)&__func);
	return __func(entityHandle, className, memberName, element);
}

static void SetEntSchemaVector4D(int32_t entityHandle, String* className, String* memberName, Vector4* value, bool changeState, int32_t element) {
	typedef void (*SetEntSchemaVector4DFn)(int32_t, String*, String*, Vector4*, bool, int32_t);
	static SetEntSchemaVector4DFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("s2sdk.SetEntSchemaVector4D", (void**)&__func);
	__func(entityHandle, className, memberName, value, changeState, element);
}

static int32_t GetEntSchemaEnt(int32_t entityHandle, String* className, String* memberName, int32_t element) {
	typedef int32_t (*GetEntSchemaEntFn)(int32_t, String*, String*, int32_t);
	static GetEntSchemaEntFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("s2sdk.GetEntSchemaEnt", (void**)&__func);
	return __func(entityHandle, className, memberName, element);
}

static void SetEntSchemaEnt(int32_t entityHandle, String* className, String* memberName, int32_t value, bool changeState, int32_t element) {
	typedef void (*SetEntSchemaEntFn)(int32_t, String*, String*, int32_t, bool, int32_t);
	static SetEntSchemaEntFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("s2sdk.SetEntSchemaEnt", (void**)&__func);
	__func(entityHandle, className, memberName, value, changeState, element);
}

static void NetworkStateChanged(int32_t entityHandle, String* className, String* memberName) {
	typedef void (*NetworkStateChangedFn)(int32_t, String*, String*);
	static NetworkStateChangedFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("s2sdk.NetworkStateChanged", (void**)&__func);
	__func(entityHandle, className, memberName);
}

static uint64_t CreateTimer(double interval, void* callback, int32_t flags, Vector* userData) {
	typedef uint64_t (*CreateTimerFn)(double, void*, int32_t, Vector*);
	static CreateTimerFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("s2sdk.CreateTimer", (void**)&__func);
	return __func(interval, callback, flags, userData);
}

static void KillsTimer(uint64_t timer) {
	typedef void (*KillsTimerFn)(uint64_t);
	static KillsTimerFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("s2sdk.KillsTimer", (void**)&__func);
	__func(timer);
}

static double GetTickInterval() {
	typedef double (*GetTickIntervalFn)();
	static GetTickIntervalFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("s2sdk.GetTickInterval", (void**)&__func);
	return __func();
}

static double GetTickedTime() {
	typedef double (*GetTickedTimeFn)();
	static GetTickedTimeFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("s2sdk.GetTickedTime", (void**)&__func);
	return __func();
}

static void OnClientConnect_Register(void* callback) {
	typedef void (*OnClientConnect_RegisterFn)(void*);
	static OnClientConnect_RegisterFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("s2sdk.OnClientConnect_Register", (void**)&__func);
	__func(callback);
}

static void OnClientConnect_Unregister(void* callback) {
	typedef void (*OnClientConnect_UnregisterFn)(void*);
	static OnClientConnect_UnregisterFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("s2sdk.OnClientConnect_Unregister", (void**)&__func);
	__func(callback);
}

static void OnClientConnect_Post_Register(void* callback) {
	typedef void (*OnClientConnect_Post_RegisterFn)(void*);
	static OnClientConnect_Post_RegisterFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("s2sdk.OnClientConnect_Post_Register", (void**)&__func);
	__func(callback);
}

static void OnClientConnect_Post_Unregister(void* callback) {
	typedef void (*OnClientConnect_Post_UnregisterFn)(void*);
	static OnClientConnect_Post_UnregisterFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("s2sdk.OnClientConnect_Post_Unregister", (void**)&__func);
	__func(callback);
}

static void OnClientConnected_Register(void* callback) {
	typedef void (*OnClientConnected_RegisterFn)(void*);
	static OnClientConnected_RegisterFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("s2sdk.OnClientConnected_Register", (void**)&__func);
	__func(callback);
}

static void OnClientConnected_Unregister(void* callback) {
	typedef void (*OnClientConnected_UnregisterFn)(void*);
	static OnClientConnected_UnregisterFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("s2sdk.OnClientConnected_Unregister", (void**)&__func);
	__func(callback);
}

static void OnClientPutInServer_Register(void* callback) {
	typedef void (*OnClientPutInServer_RegisterFn)(void*);
	static OnClientPutInServer_RegisterFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("s2sdk.OnClientPutInServer_Register", (void**)&__func);
	__func(callback);
}

static void OnClientPutInServer_Unregister(void* callback) {
	typedef void (*OnClientPutInServer_UnregisterFn)(void*);
	static OnClientPutInServer_UnregisterFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("s2sdk.OnClientPutInServer_Unregister", (void**)&__func);
	__func(callback);
}

static void OnClientDisconnect_Register(void* callback) {
	typedef void (*OnClientDisconnect_RegisterFn)(void*);
	static OnClientDisconnect_RegisterFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("s2sdk.OnClientDisconnect_Register", (void**)&__func);
	__func(callback);
}

static void OnClientDisconnect_Unregister(void* callback) {
	typedef void (*OnClientDisconnect_UnregisterFn)(void*);
	static OnClientDisconnect_UnregisterFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("s2sdk.OnClientDisconnect_Unregister", (void**)&__func);
	__func(callback);
}

static void OnClientDisconnect_Post_Register(void* callback) {
	typedef void (*OnClientDisconnect_Post_RegisterFn)(void*);
	static OnClientDisconnect_Post_RegisterFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("s2sdk.OnClientDisconnect_Post_Register", (void**)&__func);
	__func(callback);
}

static void OnClientDisconnect_Post_Unregister(void* callback) {
	typedef void (*OnClientDisconnect_Post_UnregisterFn)(void*);
	static OnClientDisconnect_Post_UnregisterFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("s2sdk.OnClientDisconnect_Post_Unregister", (void**)&__func);
	__func(callback);
}

static void OnClientActive_Register(void* callback) {
	typedef void (*OnClientActive_RegisterFn)(void*);
	static OnClientActive_RegisterFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("s2sdk.OnClientActive_Register", (void**)&__func);
	__func(callback);
}

static void OnClientActive_Unregister(void* callback) {
	typedef void (*OnClientActive_UnregisterFn)(void*);
	static OnClientActive_UnregisterFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("s2sdk.OnClientActive_Unregister", (void**)&__func);
	__func(callback);
}

static void OnClientFullyConnect_Register(void* callback) {
	typedef void (*OnClientFullyConnect_RegisterFn)(void*);
	static OnClientFullyConnect_RegisterFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("s2sdk.OnClientFullyConnect_Register", (void**)&__func);
	__func(callback);
}

static void OnClientFullyConnect_Unregister(void* callback) {
	typedef void (*OnClientFullyConnect_UnregisterFn)(void*);
	static OnClientFullyConnect_UnregisterFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("s2sdk.OnClientFullyConnect_Unregister", (void**)&__func);
	__func(callback);
}

static void OnClientSettingsChanged_Register(void* callback) {
	typedef void (*OnClientSettingsChanged_RegisterFn)(void*);
	static OnClientSettingsChanged_RegisterFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("s2sdk.OnClientSettingsChanged_Register", (void**)&__func);
	__func(callback);
}

static void OnClientSettingsChanged_Unregister(void* callback) {
	typedef void (*OnClientSettingsChanged_UnregisterFn)(void*);
	static OnClientSettingsChanged_UnregisterFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("s2sdk.OnClientSettingsChanged_Unregister", (void**)&__func);
	__func(callback);
}

static void OnClientAuthenticated_Register(void* callback) {
	typedef void (*OnClientAuthenticated_RegisterFn)(void*);
	static OnClientAuthenticated_RegisterFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("s2sdk.OnClientAuthenticated_Register", (void**)&__func);
	__func(callback);
}

static void OnClientAuthenticated_Unregister(void* callback) {
	typedef void (*OnClientAuthenticated_UnregisterFn)(void*);
	static OnClientAuthenticated_UnregisterFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("s2sdk.OnClientAuthenticated_Unregister", (void**)&__func);
	__func(callback);
}

static void OnLevelInit_Register(void* callback) {
	typedef void (*OnLevelInit_RegisterFn)(void*);
	static OnLevelInit_RegisterFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("s2sdk.OnLevelInit_Register", (void**)&__func);
	__func(callback);
}

static void OnLevelInit_Unregister(void* callback) {
	typedef void (*OnLevelInit_UnregisterFn)(void*);
	static OnLevelInit_UnregisterFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("s2sdk.OnLevelInit_Unregister", (void**)&__func);
	__func(callback);
}

static void OnLevelShutdown_Register(void* callback) {
	typedef void (*OnLevelShutdown_RegisterFn)(void*);
	static OnLevelShutdown_RegisterFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("s2sdk.OnLevelShutdown_Register", (void**)&__func);
	__func(callback);
}

static void OnLevelShutdown_Unregister(void* callback) {
	typedef void (*OnLevelShutdown_UnregisterFn)(void*);
	static OnLevelShutdown_UnregisterFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("s2sdk.OnLevelShutdown_Unregister", (void**)&__func);
	__func(callback);
}

static void OnEntitySpawned_Register(void* callback) {
	typedef void (*OnEntitySpawned_RegisterFn)(void*);
	static OnEntitySpawned_RegisterFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("s2sdk.OnEntitySpawned_Register", (void**)&__func);
	__func(callback);
}

static void OnEntitySpawned_Unregister(void* callback) {
	typedef void (*OnEntitySpawned_UnregisterFn)(void*);
	static OnEntitySpawned_UnregisterFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("s2sdk.OnEntitySpawned_Unregister", (void**)&__func);
	__func(callback);
}

static void OnEntityCreated_Register(void* callback) {
	typedef void (*OnEntityCreated_RegisterFn)(void*);
	static OnEntityCreated_RegisterFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("s2sdk.OnEntityCreated_Register", (void**)&__func);
	__func(callback);
}

static void OnEntityCreated_Unregister(void* callback) {
	typedef void (*OnEntityCreated_UnregisterFn)(void*);
	static OnEntityCreated_UnregisterFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("s2sdk.OnEntityCreated_Unregister", (void**)&__func);
	__func(callback);
}

static void OnEntityDeleted_Register(void* callback) {
	typedef void (*OnEntityDeleted_RegisterFn)(void*);
	static OnEntityDeleted_RegisterFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("s2sdk.OnEntityDeleted_Register", (void**)&__func);
	__func(callback);
}

static void OnEntityDeleted_Unregister(void* callback) {
	typedef void (*OnEntityDeleted_UnregisterFn)(void*);
	static OnEntityDeleted_UnregisterFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("s2sdk.OnEntityDeleted_Unregister", (void**)&__func);
	__func(callback);
}

static void OnEntityParentChanged_Register(void* callback) {
	typedef void (*OnEntityParentChanged_RegisterFn)(void*);
	static OnEntityParentChanged_RegisterFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("s2sdk.OnEntityParentChanged_Register", (void**)&__func);
	__func(callback);
}

static void OnEntityParentChanged_Unregister(void* callback) {
	typedef void (*OnEntityParentChanged_UnregisterFn)(void*);
	static OnEntityParentChanged_UnregisterFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("s2sdk.OnEntityParentChanged_Unregister", (void**)&__func);
	__func(callback);
}

static void OnServerStartup_Register(void* callback) {
	typedef void (*OnServerStartup_RegisterFn)(void*);
	static OnServerStartup_RegisterFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("s2sdk.OnServerStartup_Register", (void**)&__func);
	__func(callback);
}

static void OnServerStartup_Unregister(void* callback) {
	typedef void (*OnServerStartup_UnregisterFn)(void*);
	static OnServerStartup_UnregisterFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("s2sdk.OnServerStartup_Unregister", (void**)&__func);
	__func(callback);
}

static void OnServerActivate_Register(void* callback) {
	typedef void (*OnServerActivate_RegisterFn)(void*);
	static OnServerActivate_RegisterFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("s2sdk.OnServerActivate_Register", (void**)&__func);
	__func(callback);
}

static void OnServerActivate_Unregister(void* callback) {
	typedef void (*OnServerActivate_UnregisterFn)(void*);
	static OnServerActivate_UnregisterFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("s2sdk.OnServerActivate_Unregister", (void**)&__func);
	__func(callback);
}

static void OnGameFrame_Register(void* callback) {
	typedef void (*OnGameFrame_RegisterFn)(void*);
	static OnGameFrame_RegisterFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("s2sdk.OnGameFrame_Register", (void**)&__func);
	__func(callback);
}

static void OnGameFrame_Unregister(void* callback) {
	typedef void (*OnGameFrame_UnregisterFn)(void*);
	static OnGameFrame_UnregisterFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("s2sdk.OnGameFrame_Unregister", (void**)&__func);
	__func(callback);
}

static void OnUpdateWhenNotInGame_Register(void* callback) {
	typedef void (*OnUpdateWhenNotInGame_RegisterFn)(void*);
	static OnUpdateWhenNotInGame_RegisterFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("s2sdk.OnUpdateWhenNotInGame_Register", (void**)&__func);
	__func(callback);
}

static void OnUpdateWhenNotInGame_Unregister(void* callback) {
	typedef void (*OnUpdateWhenNotInGame_UnregisterFn)(void*);
	static OnUpdateWhenNotInGame_UnregisterFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("s2sdk.OnUpdateWhenNotInGame_Unregister", (void**)&__func);
	__func(callback);
}

static void OnPreWorldUpdate_Register(void* callback) {
	typedef void (*OnPreWorldUpdate_RegisterFn)(void*);
	static OnPreWorldUpdate_RegisterFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("s2sdk.OnPreWorldUpdate_Register", (void**)&__func);
	__func(callback);
}

static void OnPreWorldUpdate_Unregister(void* callback) {
	typedef void (*OnPreWorldUpdate_UnregisterFn)(void*);
	static OnPreWorldUpdate_UnregisterFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("s2sdk.OnPreWorldUpdate_Unregister", (void**)&__func);
	__func(callback);
}

static uintptr_t GetGameRules() {
	typedef uintptr_t (*GetGameRulesFn)();
	static GetGameRulesFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("s2sdk.GetGameRules", (void**)&__func);
	return __func();
}

static void TerminateRound(float delay, int32_t reason) {
	typedef void (*TerminateRoundFn)(float, int32_t);
	static TerminateRoundFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("s2sdk.TerminateRound", (void**)&__func);
	__func(delay, reason);
}

static uintptr_t GetWeaponVData(int32_t entityHandle) {
	typedef uintptr_t (*GetWeaponVDataFn)(int32_t);
	static GetWeaponVDataFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("s2sdk.GetWeaponVData", (void**)&__func);
	return __func(entityHandle);
}

static uint16_t GetWeaponDefIndex(int32_t entityHandle) {
	typedef uint16_t (*GetWeaponDefIndexFn)(int32_t);
	static GetWeaponDefIndexFn __func = NULL;
	if (__func == NULL) Plugify_GetMethodPtr2("s2sdk.GetWeaponDefIndex", (void**)&__func);
	return __func(entityHandle);
}
