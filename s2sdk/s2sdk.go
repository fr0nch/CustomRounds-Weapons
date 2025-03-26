package s2sdk

/*
#include "s2sdk.h"
#cgo noescape GetPlayerSlotFromEntityPointer
#cgo noescape GetClientBaseFromPlayerSlot
#cgo noescape GetPlayerSlotFromClientBase
#cgo noescape GetClientAuthId
#cgo noescape GetClientAccountId
#cgo noescape GetClientIp
#cgo noescape GetClientName
#cgo noescape GetClientTime
#cgo noescape GetClientLatency
#cgo noescape GetUserFlagBits
#cgo noescape SetUserFlagBits
#cgo noescape AddUserFlags
#cgo noescape RemoveUserFlags
#cgo noescape IsClientAuthorized
#cgo noescape IsClientConnected
#cgo noescape IsClientInGame
#cgo noescape IsClientSourceTV
#cgo noescape IsClientAlive
#cgo noescape IsFakeClient
#cgo noescape GetClientTeam
#cgo noescape GetClientHealth
#cgo noescape GetClientArmor
#cgo noescape GetClientAbsOrigin
#cgo noescape GetClientAbsAngles
#cgo noescape GetClientEyeAngles
#cgo noescape ProcessTargetString
#cgo noescape ChangeClientTeam
#cgo noescape SwitchClientTeam
#cgo noescape RespawnClient
#cgo noescape ForcePlayerSuicide
#cgo noescape KickClient
#cgo noescape BanClient
#cgo noescape BanIdentity
#cgo noescape GetClientActiveWeapon
#cgo noescape GetClientWeapons
#cgo noescape DropWeapon
#cgo noescape StripWeapons
#cgo noescape BumpWeapon
#cgo noescape SwitchWeapon
#cgo noescape RemoveWeapon
#cgo noescape GiveNamedItem
#cgo noescape GetClientButtons
#cgo noescape GetClientMoney
#cgo noescape SetClientMoney
#cgo noescape GetClientKills
#cgo noescape SetClientKills
#cgo noescape GetClientDeaths
#cgo noescape SetClientDeaths
#cgo noescape GetClientAssists
#cgo noescape SetClientAssists
#cgo noescape GetClientDamage
#cgo noescape SetClientDamage
#cgo noescape AddAdminCommand
#cgo noescape AddConsoleCommand
#cgo noescape RemoveCommand
#cgo noescape AddCommandListener
#cgo noescape RemoveCommandListener
#cgo noescape ServerCommand
#cgo noescape ServerCommandEx
#cgo noescape ClientCommand
#cgo noescape FakeClientCommand
#cgo noescape PrintToServer
#cgo noescape PrintToConsole
#cgo noescape PrintToChat
#cgo noescape PrintCenterText
#cgo noescape PrintAlertText
#cgo noescape PrintCentreHtml
#cgo noescape PrintToConsoleAll
#cgo noescape PrintToChatAll
#cgo noescape PrintCenterTextAll
#cgo noescape PrintAlertTextAll
#cgo noescape PrintCentreHtmlAll
#cgo noescape PrintToChatColored
#cgo noescape PrintToChatColoredAll
#cgo noescape CreateConVar
#cgo noescape CreateConVarBool
#cgo noescape CreateConVarInt16
#cgo noescape CreateConVarUInt16
#cgo noescape CreateConVarInt32
#cgo noescape CreateConVarUInt32
#cgo noescape CreateConVarInt64
#cgo noescape CreateConVarUInt64
#cgo noescape CreateConVarFloat
#cgo noescape CreateConVarDouble
#cgo noescape CreateConVarColor
#cgo noescape CreateConVarVector2
#cgo noescape CreateConVarVector3
#cgo noescape CreateConVarVector4
#cgo noescape CreateConVarQAngle
#cgo noescape FindConVar
#cgo noescape FindConVar2
#cgo noescape HookConVarChange
#cgo noescape UnhookConVarChange
#cgo noescape IsConVarFlagSet
#cgo noescape AddConVarFlags
#cgo noescape RemoveConVarFlags
#cgo noescape GetConVarFlags
#cgo noescape GetConVarBounds
#cgo noescape SetConVarBounds
#cgo noescape GetConVarDefault
#cgo noescape GetConVarValue
#cgo noescape GetConVar
#cgo noescape GetConVarBool
#cgo noescape GetConVarInt16
#cgo noescape GetConVarUInt16
#cgo noescape GetConVarInt32
#cgo noescape GetConVarUInt32
#cgo noescape GetConVarInt64
#cgo noescape GetConVarUInt64
#cgo noescape GetConVarFloat
#cgo noescape GetConVarDouble
#cgo noescape GetConVarString
#cgo noescape GetConVarColor
#cgo noescape GetConVarVector2
#cgo noescape GetConVarVector
#cgo noescape GetConVarVector4
#cgo noescape GetConVarQAngle
#cgo noescape SetConVarValue
#cgo noescape SetConVar
#cgo noescape SetConVarBool
#cgo noescape SetConVarInt16
#cgo noescape SetConVarUInt16
#cgo noescape SetConVarInt32
#cgo noescape SetConVarUInt32
#cgo noescape SetConVarInt64
#cgo noescape SetConVarUInt64
#cgo noescape SetConVarFloat
#cgo noescape SetConVarDouble
#cgo noescape SetConVarString
#cgo noescape SetConVarColor
#cgo noescape SetConVarVector2
#cgo noescape SetConVarVector3
#cgo noescape SetConVarVector4
#cgo noescape SetConVarQAngle
#cgo noescape SendConVarValue
#cgo noescape GetClientConVarValue
#cgo noescape SetFakeClientConVarValue
#cgo noescape GetGameDirectory
#cgo noescape GetCurrentMap
#cgo noescape IsMapValid
#cgo noescape GetGameTime
#cgo noescape GetGameTickCount
#cgo noescape GetGameFrameTime
#cgo noescape GetEngineTime
#cgo noescape GetMaxClients
#cgo noescape PrecacheGeneric
#cgo noescape IsGenericPrecache
#cgo noescape PrecacheModel
#cgo noescape IsModelPrecache
#cgo noescape PrecacheSound
#cgo noescape IsSoundPrecached
#cgo noescape PrecacheDecal
#cgo noescape IsDecalPrecached
#cgo noescape GetEconItemSystem
#cgo noescape IsServerPaused
#cgo noescape QueueTaskForNextFrame
#cgo noescape QueueTaskForNextWorldUpdate
#cgo noescape GetSoundDuration
#cgo noescape EmitSound
#cgo noescape EmitSoundToClient
#cgo noescape EntIndexToEntPointer
#cgo noescape EntPointerToEntIndex
#cgo noescape EntPointerToEntHandle
#cgo noescape EntHandleToEntPointer
#cgo noescape EntIndexToEntHandle
#cgo noescape EntHandleToEntIndex
#cgo noescape IsValidEntHandle
#cgo noescape IsValidEntPointer
#cgo noescape GetFirstActiveEntity
#cgo noescape GetConcreteEntityListPointer
#cgo noescape HookEntityOutput
#cgo noescape UnhookEntityOutput
#cgo noescape FindEntityByClassname
#cgo noescape FindEntityByName
#cgo noescape CreateEntityByName
#cgo noescape DispatchSpawn
#cgo noescape DispatchSpawn2
#cgo noescape RemoveEntity
#cgo noescape GetEntityClassname
#cgo noescape GetEntityName
#cgo noescape SetEntityName
#cgo noescape GetEntityMoveType
#cgo noescape SetEntityMoveType
#cgo noescape GetEntityGravity
#cgo noescape SetEntityGravity
#cgo noescape GetEntityFlags
#cgo noescape SetEntityFlags
#cgo noescape GetEntityRenderColor
#cgo noescape SetEntityRenderColor
#cgo noescape GetEntityRenderMode
#cgo noescape SetEntityRenderMode
#cgo noescape GetEntityHealth
#cgo noescape SetEntityHealth
#cgo noescape GetTeamEntity
#cgo noescape SetTeamEntity
#cgo noescape GetEntityOwner
#cgo noescape SetEntityOwner
#cgo noescape GetEntityParent
#cgo noescape SetEntityParent
#cgo noescape GetEntityAbsOrigin
#cgo noescape SetEntityAbsOrigin
#cgo noescape GetEntityAngRotation
#cgo noescape SetEntityAngRotation
#cgo noescape GetEntityAbsVelocity
#cgo noescape SetEntityAbsVelocity
#cgo noescape GetEntityModel
#cgo noescape SetEntityModel
#cgo noescape GetEntityWaterLevel
#cgo noescape GetEntityGroundEntity
#cgo noescape GetEntityEffects
#cgo noescape TeleportEntity
#cgo noescape AcceptInput
#cgo noescape HookEvent
#cgo noescape UnhookEvent
#cgo noescape CreateEvent
#cgo noescape FireEvent
#cgo noescape FireEventToClient
#cgo noescape CancelCreatedEvent
#cgo noescape GetEventBool
#cgo noescape GetEventFloat
#cgo noescape GetEventInt
#cgo noescape GetEventUInt64
#cgo noescape GetEventString
#cgo noescape GetEventPtr
#cgo noescape GetEventPlayerController
#cgo noescape GetEventPlayerIndex
#cgo noescape GetEventPlayerPawn
#cgo noescape GetEventEntity
#cgo noescape GetEventEntityIndex
#cgo noescape GetEventEntityHandle
#cgo noescape GetEventName
#cgo noescape SetEventBool
#cgo noescape SetEventFloat
#cgo noescape SetEventInt
#cgo noescape SetEventUInt64
#cgo noescape SetEventString
#cgo noescape SetEventPtr
#cgo noescape SetEventPlayerController
#cgo noescape SetEventPlayerIndex
#cgo noescape SetEventEntity
#cgo noescape SetEventEntityIndex
#cgo noescape SetEventEntityHandle
#cgo noescape SetEventBroadcast
#cgo noescape LoadEventsFromFile
#cgo noescape CloseGameConfigFile
#cgo noescape LoadGameConfigFile
#cgo noescape GetGameConfigPath
#cgo noescape GetGameConfigLibrary
#cgo noescape GetGameConfigSignature
#cgo noescape GetGameConfigSymbol
#cgo noescape GetGameConfigPatch
#cgo noescape GetGameConfigOffset
#cgo noescape GetGameConfigAddress
#cgo noescape GetGameConfigMemSig
#cgo noescape RegisterLoggingChannel
#cgo noescape AddLoggerTagToChannel
#cgo noescape HasLoggerTag
#cgo noescape IsLoggerChannelEnabledBySeverity
#cgo noescape IsLoggerChannelEnabledByVerbosity
#cgo noescape GetLoggerChannelVerbosity
#cgo noescape SetLoggerChannelVerbosity
#cgo noescape SetLoggerChannelVerbosityByName
#cgo noescape SetLoggerChannelVerbosityByTag
#cgo noescape GetLoggerChannelColor
#cgo noescape SetLoggerChannelColor
#cgo noescape GetLoggerChannelFlags
#cgo noescape SetLoggerChannelFlags
#cgo noescape Log
#cgo noescape LogColored
#cgo noescape LogFull
#cgo noescape LogFullColored
#cgo noescape GetSchemaOffset
#cgo noescape GetSchemaChainOffset
#cgo noescape IsSchemaFieldNetworked
#cgo noescape GetSchemaClassSize
#cgo noescape GetEntData2
#cgo noescape SetEntData2
#cgo noescape GetEntDataFloat2
#cgo noescape SetEntDataFloat2
#cgo noescape GetEntDataString2
#cgo noescape SetEntDataString2
#cgo noescape GetEntDataVector2
#cgo noescape SetEntDataVector2
#cgo noescape GetEntDataEnt2
#cgo noescape SetEntDataEnt2
#cgo noescape ChangeEntityState2
#cgo noescape GetEntData
#cgo noescape SetEntData
#cgo noescape GetEntDataFloat
#cgo noescape SetEntDataFloat
#cgo noescape GetEntDataString
#cgo noescape SetEntDataString
#cgo noescape GetEntDataVector
#cgo noescape SetEntDataVector
#cgo noescape GetEntDataEnt
#cgo noescape SetEntDataEnt
#cgo noescape ChangeEntityState
#cgo noescape GetEntSchemaArraySize2
#cgo noescape GetEntSchema2
#cgo noescape SetEntSchema2
#cgo noescape GetEntSchemaFloat2
#cgo noescape SetEntSchemaFloat2
#cgo noescape GetEntSchemaString2
#cgo noescape SetEntSchemaString2
#cgo noescape GetEntSchemaVector3D2
#cgo noescape SetEntSchemaVector3D2
#cgo noescape GetEntSchemaVector2D2
#cgo noescape SetEntSchemaVector2D2
#cgo noescape GetEntSchemaVector4D2
#cgo noescape SetEntSchemaVector4D2
#cgo noescape GetEntSchemaEnt2
#cgo noescape SetEntSchemaEnt2
#cgo noescape NetworkStateChanged2
#cgo noescape GetEntSchemaArraySize
#cgo noescape GetEntSchema
#cgo noescape SetEntSchema
#cgo noescape GetEntSchemaFloat
#cgo noescape SetEntSchemaFloat
#cgo noescape GetEntSchemaString
#cgo noescape SetEntSchemaString
#cgo noescape GetEntSchemaVector3D
#cgo noescape SetEntSchemaVector3D
#cgo noescape GetEntSchemaVector2D
#cgo noescape SetEntSchemaVector2D
#cgo noescape GetEntSchemaVector4D
#cgo noescape SetEntSchemaVector4D
#cgo noescape GetEntSchemaEnt
#cgo noescape SetEntSchemaEnt
#cgo noescape NetworkStateChanged
#cgo noescape CreateTimer
#cgo noescape KillsTimer
#cgo noescape GetTickInterval
#cgo noescape GetTickedTime
#cgo noescape OnClientConnect_Register
#cgo noescape OnClientConnect_Unregister
#cgo noescape OnClientConnect_Post_Register
#cgo noescape OnClientConnect_Post_Unregister
#cgo noescape OnClientConnected_Register
#cgo noescape OnClientConnected_Unregister
#cgo noescape OnClientPutInServer_Register
#cgo noescape OnClientPutInServer_Unregister
#cgo noescape OnClientDisconnect_Register
#cgo noescape OnClientDisconnect_Unregister
#cgo noescape OnClientDisconnect_Post_Register
#cgo noescape OnClientDisconnect_Post_Unregister
#cgo noescape OnClientActive_Register
#cgo noescape OnClientActive_Unregister
#cgo noescape OnClientFullyConnect_Register
#cgo noescape OnClientFullyConnect_Unregister
#cgo noescape OnClientSettingsChanged_Register
#cgo noescape OnClientSettingsChanged_Unregister
#cgo noescape OnClientAuthenticated_Register
#cgo noescape OnClientAuthenticated_Unregister
#cgo noescape OnLevelInit_Register
#cgo noescape OnLevelInit_Unregister
#cgo noescape OnLevelShutdown_Register
#cgo noescape OnLevelShutdown_Unregister
#cgo noescape OnEntitySpawned_Register
#cgo noescape OnEntitySpawned_Unregister
#cgo noescape OnEntityCreated_Register
#cgo noescape OnEntityCreated_Unregister
#cgo noescape OnEntityDeleted_Register
#cgo noescape OnEntityDeleted_Unregister
#cgo noescape OnEntityParentChanged_Register
#cgo noescape OnEntityParentChanged_Unregister
#cgo noescape OnServerStartup_Register
#cgo noescape OnServerStartup_Unregister
#cgo noescape OnServerActivate_Register
#cgo noescape OnServerActivate_Unregister
#cgo noescape OnGameFrame_Register
#cgo noescape OnGameFrame_Unregister
#cgo noescape OnUpdateWhenNotInGame_Register
#cgo noescape OnUpdateWhenNotInGame_Unregister
#cgo noescape OnPreWorldUpdate_Register
#cgo noescape OnPreWorldUpdate_Unregister
#cgo noescape GetGameRules
#cgo noescape TerminateRound
#cgo noescape GetWeaponVData
#cgo noescape GetWeaponDefIndex
*/
import "C"
import (
	"github.com/untrustedmodders/go-plugify"
	"unsafe"
)

// Generated with https://github.com/untrustedmodders/plugify-module-golang/blob/main/generator/generator.py from s2sdk

// CSTeam - Enum representing the possible teams in Counter-Strike.
type CSTeam = int32

const (
	// None - No team.
	None__ CSTeam = 0
	// Spectator - Spectator team.
	Spectator CSTeam = 1
	// T - Terrorist team.
	T CSTeam = 2
	// CT - Counter-Terrorist team.
	CT CSTeam = 3
)

// ConVarFlag - Enum representing various flags for ConVars and ConCommands.
type ConVarFlag = int64

const (
	// None - The default, no flags at all.
	None___ ConVarFlag = 0
	// LinkedConcommand - Linked to a ConCommand.
	LinkedConcommand ConVarFlag = 1
	// DevelopmentOnly - Hidden in released products. Automatically removed if ALLOW_DEVELOPMENT_CVARS is defined.
	DevelopmentOnly ConVarFlag = 2
	// GameDll - Defined by the game DLL.
	GameDll ConVarFlag = 4
	// ClientDll - Defined by the client DLL.
	ClientDll ConVarFlag = 8
	// Hidden - Hidden. Doesn't appear in find or auto-complete. Like DEVELOPMENTONLY but cannot be compiled out.
	Hidden ConVarFlag = 16
	// Protected - Server cvar; data is not sent since it's sensitive (e.g., passwords).
	Protected ConVarFlag = 32
	// SpOnly - This cvar cannot be changed by clients connected to a multiplayer server.
	SpOnly ConVarFlag = 64
	// Archive - Saved to vars.rc.
	Archive ConVarFlag = 128
	// Notify - Notifies players when changed.
	Notify ConVarFlag = 256
	// UserInfo - Changes the client's info string.
	UserInfo ConVarFlag = 512
	// Missing0 - Hides the cvar from lookups.
	Missing0 ConVarFlag = 1024
	// Unlogged - If this is a server cvar, changes are not logged to the file or console.
	Unlogged ConVarFlag = 2048
	// Missing1 - Hides the cvar from lookups.
	Missing1 ConVarFlag = 4096
	// Replicated - Server-enforced setting on clients.
	Replicated ConVarFlag = 8192
	// Cheat - Only usable in singleplayer/debug or multiplayer with sv_cheats.
	Cheat ConVarFlag = 16384
	// PerUser - Causes auto-generated varnameN for splitscreen slots.
	PerUser ConVarFlag = 32768
	// Demo - Records this cvar when starting a demo file.
	Demo ConVarFlag = 65536
	// DontRecord - Excluded from demo files.
	DontRecord ConVarFlag = 131072
	// Missing2 - Reserved for future use.
	Missing2 ConVarFlag = 262144
	// Release - Cvars tagged with this are available to customers.
	Release ConVarFlag = 524288
	// MenuBarItem - Marks the cvar as a menu bar item.
	MenuBarItem ConVarFlag = 1048576
	// Missing3 - Reserved for future use.
	Missing3 ConVarFlag = 2097152
	// NotConnected - Cannot be changed by a client connected to a server.
	NotConnected ConVarFlag = 4194304
	// VconsoleFuzzyMatching - Enables fuzzy matching for vconsole.
	VconsoleFuzzyMatching ConVarFlag = 8388608
	// ServerCanExecute - The server can execute this command on clients.
	ServerCanExecute ConVarFlag = 16777216
	// ClientCanExecute - Allows clients to execute this command.
	ClientCanExecute ConVarFlag = 33554432
	// ServerCannotQuery - The server cannot query this cvar's value.
	ServerCannotQuery ConVarFlag = 67108864
	// VconsoleSetFocus - Sets focus in the vconsole.
	VconsoleSetFocus ConVarFlag = 134217728
	// ClientCmdCanExecute - IVEngineClient::ClientCmd can execute this command.
	ClientCmdCanExecute ConVarFlag = 268435456
	// ExecutePerTick - Executes the cvar every tick.
	ExecutePerTick ConVarFlag = 536870912
)

// ResultType - Enum representing the possible results of an operation.
type ResultType = int32

const (
	// Continue - The action continues to be processed without interruption.
	Continue ResultType = 0
	// Changed - Indicates that the action has altered the state or behavior during execution.
	Changed ResultType = 1
	// Handled - The action has been successfully handled, and no further action is required.
	Handled ResultType = 2
	// Stop - The action processing is halted, and no further steps will be executed.
	Stop ResultType = 3
)

// HookMode - Enum representing the type of callback.
type HookMode = uint8

const (
	// Pre - Callback will be executed before the original function
	Pre HookMode = 0
	// Post - Callback will be executed after the original function
	Post HookMode = 1
)

type ConVarType = int16

const (
	// Invalid - Invalid type
	Invalid ConVarType = -1
	// Bool - Boolean type
	Bool ConVarType = 0
	// Int16 - 16-bit signed integer
	Int16 ConVarType = 1
	// UInt16 - 16-bit unsigned integer
	UInt16 ConVarType = 2
	// Int32 - 32-bit signed integer
	Int32_ ConVarType = 3
	// UInt32 - 32-bit unsigned integer
	UInt32_ ConVarType = 4
	// Int64 - 64-bit signed integer
	Int64_ ConVarType = 5
	// UInt64 - 64-bit unsigned integer
	UInt64_ ConVarType = 6
	// Float32 - 32-bit floating point
	Float32__ ConVarType = 7
	// Float64 - 64-bit floating point (double)
	Float64_ ConVarType = 8
	// String - String type
	String_ ConVarType = 9
	// Color - Color type
	Color ConVarType = 10
	// Vector2 - 2D vector
	Vector2 ConVarType = 11
	// Vector3 - 3D vector
	Vector3 ConVarType = 12
	// Vector4 - 4D vector
	Vector4 ConVarType = 13
	// Qangle - Quaternion angle
	Qangle ConVarType = 14
	// Max - Maximum value (used for bounds checking)
	Max_ ConVarType = 15
)

// MoveType - Enum representing various movement types for entities.
type MoveType = int32

const (
	// None - Never moves.
	None_ MoveType = 0
	// Isometric - Previously isometric movement type.
	Isometric MoveType = 1
	// Walk - Player only - moving on the ground.
	Walk MoveType = 2
	// Fly - No gravity, but still collides with stuff.
	Fly MoveType = 3
	// Flygravity - Flies through the air and is affected by gravity.
	Flygravity MoveType = 4
	// Vphysics - Uses VPHYSICS for simulation.
	Vphysics MoveType = 5
	// Push - No clip to world, push and crush.
	Push MoveType = 6
	// Noclip - No gravity, no collisions, still has velocity/avelocity.
	Noclip MoveType = 7
	// Ladder - Used by players only when going onto a ladder.
	Ladder MoveType = 8
	// Observer - Observer movement, depends on player's observer mode.
	Observer MoveType = 9
	// Custom - Allows the entity to describe its own physics.
	Custom MoveType = 10
)

// RenderMode - Enum representing rendering modes for materials.
type RenderMode = int8

const (
	// Normal - Standard rendering mode (src).
	Normal RenderMode = 0
	// TransColor - Composite: c*a + dest*(1-a).
	TransColor RenderMode = 1
	// TransTexture - Composite: src*a + dest*(1-a).
	TransTexture RenderMode = 2
	// Glow - Composite: src*a + dest -- No Z buffer checks -- Fixed size in screen space.
	Glow RenderMode = 3
	// TransAlpha - Composite: src*srca + dest*(1-srca).
	TransAlpha RenderMode = 4
	// TransAdd - Composite: src*a + dest.
	TransAdd RenderMode = 5
	// Environmental - Not drawn, used for environmental effects.
	Environmental RenderMode = 6
	// TransAddFrameBlend - Uses a fractional frame value to blend between animation frames.
	TransAddFrameBlend RenderMode = 7
	// TransAlphaAdd - Composite: src + dest*(1-a).
	TransAlphaAdd RenderMode = 8
	// WorldGlow - Same as Glow but not fixed size in screen space.
	WorldGlow RenderMode = 9
	// None - No rendering.
	None RenderMode = 10
	// DevVisualizer - Developer visualizer rendering mode.
	DevVisualizer RenderMode = 11
)

// FieldType - Represents the possible types of data that can be passed as a value in input actions.
type FieldType = int32

const (
	// Auto - Automatically detect the type of the value.
	Auto FieldType = 0
	// Float32 - A 32-bit floating-point number.
	Float32___ FieldType = 1
	// Float64 - A 64-bit floating-point number.
	Float64 FieldType = 2
	// Int32 - A 32-bit signed integer.
	Int32 FieldType = 3
	// UInt32 - A 32-bit unsigned integer.
	UInt32 FieldType = 4
	// Int64 - A 64-bit signed integer.
	Int64 FieldType = 5
	// UInt64 - A 64-bit unsigned integer.
	UInt64 FieldType = 6
	// Boolean - A boolean value (true or false).
	Boolean FieldType = 7
	// Character - A single character.
	Character FieldType = 8
	// String - A managed string object.
	String FieldType = 9
	// CString - A null-terminated C-style string.
	CString FieldType = 10
	// HScript - A script handle, typically for scripting integration.
	HScript FieldType = 11
	// EHandle - An entity handle, used to reference an entity within the system.
	EHandle FieldType = 12
	// Resource - A resource handle, such as a file or asset reference.
	Resource FieldType = 13
	// Vector3d - A 3D vector, typically representing position or direction.
	Vector3d FieldType = 14
	// Vector2d - A 2D vector, for planar data or coordinates.
	Vector2d FieldType = 15
	// Vector4d - A 4D vector, often used for advanced mathematical representations.
	Vector4d FieldType = 16
	// Color32 - A 32-bit color value (RGBA).
	Color32 FieldType = 17
	// QAngle - A quaternion-based angle representation.
	QAngle FieldType = 18
	// Quaternion - A quaternion, used for rotation and orientation calculations.
	Quaternion FieldType = 19
)

// LoggingVerbosity - Enum representing the possible verbosity of a logger.
type LoggingVerbosity = int32

const (
	// Off - Turns off all spew.
	Off_ LoggingVerbosity = 0
	// Essential - Turns on vital logs.
	Essential LoggingVerbosity = 1
	// Default - Turns on most messages.
	Default_ LoggingVerbosity = 2
	// Detailed - Allows for walls of text that are usually useful.
	Detailed_ LoggingVerbosity = 3
	// Max - Allows everything.
	Max LoggingVerbosity = 4
)

// LoggingSeverity - Enum representing the possible verbosity of a logger.
type LoggingSeverity = int32

const (
	// Off - Turns off all spew.
	Off__ LoggingSeverity = 0
	// Detailed - A debug message.
	Detailed__ LoggingSeverity = 1
	// Message - An informative logging message.
	Message LoggingSeverity = 2
	// Warning - A warning, typically non-fatal.
	Warning LoggingSeverity = 3
	// Assert - A message caused by an Assert**() operation.
	Assert LoggingSeverity = 4
	// Error - An error, typically fatal/unrecoverable.
	Error LoggingSeverity = 5
)

// LoggingChannelFlags - Logging channel behavior flags, set on channel creation.
type LoggingChannelFlags = int32

const (
	// ConsoleOnly - Indicates that the spew is only relevant to interactive consoles.
	ConsoleOnly LoggingChannelFlags = 1
	// DoNotEcho - Indicates that spew should not be echoed to any output devices.
	DoNotEcho LoggingChannelFlags = 2
)

// TimerFlag - Enum representing the possible flags of a timer.
type TimerFlag = int32

const (
	// Default - Timer with no unique properties.
	Default TimerFlag = 0
	// Repeat - Timer will repeat until stopped.
	Repeat TimerFlag = 1
	// NoMapChange - Timer will not carry over mapchanges.
	NoMapChange TimerFlag = 2
)

// CSRoundEndReason - Enum representing the possible reasons for a round ending in Counter-Strike.
type CSRoundEndReason = int32

const (
	// TargetBombed - Target successfully bombed.
	TargetBombed CSRoundEndReason = 1
	// VIPEscaped - The VIP has escaped (not present in CS:GO).
	VIPEscaped CSRoundEndReason = 2
	// VIPKilled - VIP has been assassinated (not present in CS:GO).
	VIPKilled CSRoundEndReason = 3
	// TerroristsEscaped - The terrorists have escaped.
	TerroristsEscaped CSRoundEndReason = 4
	// CTStoppedEscape - The CTs have prevented most of the terrorists from escaping.
	CTStoppedEscape CSRoundEndReason = 5
	// TerroristsStopped - Escaping terrorists have all been neutralized.
	TerroristsStopped CSRoundEndReason = 6
	// BombDefused - The bomb has been defused.
	BombDefused CSRoundEndReason = 7
	// CTWin - Counter-Terrorists win.
	CTWin CSRoundEndReason = 8
	// TerroristWin - Terrorists win.
	TerroristWin CSRoundEndReason = 9
	// Draw - Round draw.
	Draw CSRoundEndReason = 10
	// HostagesRescued - All hostages have been rescued.
	HostagesRescued CSRoundEndReason = 11
	// TargetSaved - Target has been saved.
	TargetSaved CSRoundEndReason = 12
	// HostagesNotRescued - Hostages have not been rescued.
	HostagesNotRescued CSRoundEndReason = 13
	// TerroristsNotEscaped - Terrorists have not escaped.
	TerroristsNotEscaped CSRoundEndReason = 14
	// VIPNotEscaped - VIP has not escaped (not present in CS:GO).
	VIPNotEscaped CSRoundEndReason = 15
	// GameStart - Game commencing.
	GameStart CSRoundEndReason = 16
	// TerroristsSurrender - Terrorists surrender.
	TerroristsSurrender CSRoundEndReason = 17
	// CTSurrender - CTs surrender.
	CTSurrender CSRoundEndReason = 18
	// TerroristsPlanted - Terrorists planted the bomb.
	TerroristsPlanted CSRoundEndReason = 19
	// CTsReachedHostage - CTs reached the hostage.
	CTsReachedHostage CSRoundEndReason = 20
	// SurvivalWin - Survival mode win.
	SurvivalWin CSRoundEndReason = 21
	// SurvivalDraw - Survival mode draw.
	SurvivalDraw CSRoundEndReason = 22
)

// CommandCallback - Handles the execution of a command triggered by a caller. This function processes the command, interprets its context, and handles any provided arguments.
type CommandCallback func(caller int32, context int32, arguments []string) ResultType

// ChangeCallback - Handles changes to a console variable's value. This function is called whenever the value of a specific console variable is modified.
type ChangeCallback func(conVarHandle uint64, newValue string, oldValue string)

// TaskCallback - Defines a QueueTask Callback.
type TaskCallback func(userData []interface{})

// HookEntityOutputCallback - This function is a callback handler for entity output events. It is triggered when a specific output event is activated, and it handles the process by passing the activator, the caller, and a delay parameter for the output.
type HookEntityOutputCallback func(activatorHandle int32, callerHandle int32, flDelay float32) ResultType

// EventCallback - Handles events triggered by the game event system. This function processes the event data, determines the necessary action, and optionally prevents event broadcasting.
type EventCallback func(name string, event uintptr, dontBroadcast bool) ResultType

// TimerCallback - This function is invoked when a timer event occurs. It handles the timer-related logic and performs necessary actions based on the event.
type TimerCallback func(timer uint64, userData []interface{})

// OnClientConnectCallback - Called on client connection. If you return true, the client will be allowed in the server. If you return false (or return nothing), the client will be rejected. If the client is rejected by this forward or any other, OnClientDisconnect will not be called.<br>Note: Do not write to rejectmsg if you plan on returning true. If multiple plugins write to the string buffer, it is not defined which plugin's string will be shown to the client, but it is guaranteed one of them will.
type OnClientConnectCallback func(playerSlot int32, name string, networkId string) bool

// OnClientConnect_PostCallback - Called on client connection.
type OnClientConnect_PostCallback func(playerSlot int32)

// OnClientConnectedCallback - Called once a client successfully connects. This callback is paired with OnClientDisconnect.
type OnClientConnectedCallback func(playerSlot int32)

// OnClientPutInServerCallback - Called when a client is entering the game.
type OnClientPutInServerCallback func(playerSlot int32)

// OnClientDisconnectCallback - Called when a client is disconnecting from the server.
type OnClientDisconnectCallback func(playerSlot int32)

// OnClientDisconnect_PostCallback - Called when a client is disconnected from the server.
type OnClientDisconnect_PostCallback func(playerSlot int32, reason int32)

// OnClientActiveCallback - Called when a client is activated by the game.
type OnClientActiveCallback func(playerSlot int32, isActive bool)

// OnClientFullyConnectCallback - Called when a client is fully connected to the game.
type OnClientFullyConnectCallback func(playerSlot int32)

// OnClientSettingsChangedCallback - Called whenever the client's settings are changed.
type OnClientSettingsChangedCallback func(playerSlot int32)

// OnClientAuthenticatedCallback - Called when a client is fully connected to the game.
type OnClientAuthenticatedCallback func(playerSlot int32, steam uint64)

// OnLevelInitCallback - Called when the map starts loading.
type OnLevelInitCallback func(mapName string, mapEntities string)

// OnLevelShutdownCallback - Called right before a map ends.
type OnLevelShutdownCallback func()

// OnEntitySpawnedCallback - Called when an entity is spawned.
type OnEntitySpawnedCallback func(entityHandle int32)

// OnEntityCreatedCallback - Called when an entity is created.
type OnEntityCreatedCallback func(entityHandle int32)

// OnEntityDeletedCallback - Called when when an entity is destroyed.
type OnEntityDeletedCallback func(entityHandle int32)

// OnEntityParentChangedCallback - When an entity is reparented to another entity.
type OnEntityParentChangedCallback func(entityHandle int32, parentHandle int32)

// OnServerStartupCallback - Called on every server startup.
type OnServerStartupCallback func()

// OnServerActivateCallback - Called on every server activate.
type OnServerActivateCallback func()

// OnGameFrameCallback - Called before every server frame. Note that you should avoid doing expensive computations or declaring large local arrays.
type OnGameFrameCallback func(simulating bool, firstTick bool, lastTick bool)
type OnUpdateWhenNotInGameCallback func(deltaTime float32)
type OnPreWorldUpdateCallback func(simulating bool)

// GetPlayerSlotFromEntityPointer - Retrieves The player slot from a given entity pointer.
// @param entity: A pointer to the entity (CBaseEntity*).
// @return The player slot if valid, otherwise -1.
func GetPlayerSlotFromEntityPointer(entity uintptr) int32 {
	var __retVal int32
	__entity := C.uintptr_t(entity)
	__retVal = int32(C.GetPlayerSlotFromEntityPointer(__entity))
	return __retVal
}

// GetClientBaseFromPlayerSlot - Retrieves the client object from a given player slot.
// @param playerSlot: The index of the client.
// @return A pointer to the client object if found, otherwise nullptr.
func GetClientBaseFromPlayerSlot(playerSlot int32) uintptr {
	var __retVal uintptr
	__playerSlot := C.int32_t(playerSlot)
	__retVal = uintptr(C.GetClientBaseFromPlayerSlot(__playerSlot))
	return __retVal
}

// GetPlayerSlotFromClientBase - Retrieves the index of a given client object.
// @param client: A pointer to the client object (CServerSideClient*).
// @return The player slot if found, otherwise -1.
func GetPlayerSlotFromClientBase(client uintptr) int32 {
	var __retVal int32
	__client := C.uintptr_t(client)
	__retVal = int32(C.GetPlayerSlotFromClientBase(__client))
	return __retVal
}

// GetClientAuthId - Retrieves a client's authentication string (SteamID).
// @param playerSlot: The index of the player's slot whose authentication string is being retrieved.
// @return The authentication string.
func GetClientAuthId(playerSlot int32) string {
	var __retVal string
	var __retVal_native plugify.PlgString
	__playerSlot := C.int32_t(playerSlot)
	plugify.Block{
		Try: func() {
			__native := C.GetClientAuthId(__playerSlot)
			__retVal_native = *(*plugify.PlgString)(unsafe.Pointer(&__native))
			// Unmarshal - Convert native data to managed data.
			__retVal = plugify.GetStringData(&__retVal_native)
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__retVal_native)
		},
	}.Do()
	return __retVal
}

// GetClientAccountId - Returns the client's Steam account ID, a unique number identifying a given Steam account.
// @param playerSlot: The index of the player's slot.
// @return Steam account ID.
func GetClientAccountId(playerSlot int32) uint64 {
	var __retVal uint64
	__playerSlot := C.int32_t(playerSlot)
	__retVal = uint64(C.GetClientAccountId(__playerSlot))
	return __retVal
}

// GetClientIp - Retrieves a client's IP address.
// @param playerSlot: The index of the player's slot.
// @return The IP address.
func GetClientIp(playerSlot int32) string {
	var __retVal string
	var __retVal_native plugify.PlgString
	__playerSlot := C.int32_t(playerSlot)
	plugify.Block{
		Try: func() {
			__native := C.GetClientIp(__playerSlot)
			__retVal_native = *(*plugify.PlgString)(unsafe.Pointer(&__native))
			// Unmarshal - Convert native data to managed data.
			__retVal = plugify.GetStringData(&__retVal_native)
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__retVal_native)
		},
	}.Do()
	return __retVal
}

// GetClientName - Returns the client's name.
// @param playerSlot: The index of the player's slot.
// @return The client's name.
func GetClientName(playerSlot int32) string {
	var __retVal string
	var __retVal_native plugify.PlgString
	__playerSlot := C.int32_t(playerSlot)
	plugify.Block{
		Try: func() {
			__native := C.GetClientName(__playerSlot)
			__retVal_native = *(*plugify.PlgString)(unsafe.Pointer(&__native))
			// Unmarshal - Convert native data to managed data.
			__retVal = plugify.GetStringData(&__retVal_native)
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__retVal_native)
		},
	}.Do()
	return __retVal
}

// GetClientTime - Returns the client's connection time in seconds.
// @param playerSlot: The index of the player's slot.
// @return float Connection time in seconds.
func GetClientTime(playerSlot int32) float32 {
	var __retVal float32
	__playerSlot := C.int32_t(playerSlot)
	__retVal = float32(C.GetClientTime(__playerSlot))
	return __retVal
}

// GetClientLatency - Returns the client's current latency (RTT).
// @param playerSlot: The index of the player's slot.
// @return float Latency value.
func GetClientLatency(playerSlot int32) float32 {
	var __retVal float32
	__playerSlot := C.int32_t(playerSlot)
	__retVal = float32(C.GetClientLatency(__playerSlot))
	return __retVal
}

// GetUserFlagBits - Returns the client's access flags.
// @param playerSlot: The index of the player's slot.
// @return uint64 Access flags as a bitmask.
func GetUserFlagBits(playerSlot int32) uint64 {
	var __retVal uint64
	__playerSlot := C.int32_t(playerSlot)
	__retVal = uint64(C.GetUserFlagBits(__playerSlot))
	return __retVal
}

// SetUserFlagBits - Sets the access flags on a client using a bitmask.
// @param playerSlot: The index of the player's slot.
// @param flags: Bitmask representing the flags to be set.
func SetUserFlagBits(playerSlot int32, flags uint64) {
	__playerSlot := C.int32_t(playerSlot)
	__flags := C.uint64_t(flags)
	C.SetUserFlagBits(__playerSlot, __flags)
}

// AddUserFlags - Adds access flags to a client.
// @param playerSlot: The index of the player's slot.
// @param flags: Bitmask representing the flags to be added.
func AddUserFlags(playerSlot int32, flags uint64) {
	__playerSlot := C.int32_t(playerSlot)
	__flags := C.uint64_t(flags)
	C.AddUserFlags(__playerSlot, __flags)
}

// RemoveUserFlags - Removes access flags from a client.
// @param playerSlot: The index of the player's slot.
// @param flags: Bitmask representing the flags to be removed.
func RemoveUserFlags(playerSlot int32, flags uint64) {
	__playerSlot := C.int32_t(playerSlot)
	__flags := C.uint64_t(flags)
	C.RemoveUserFlags(__playerSlot, __flags)
}

// IsClientAuthorized - Checks if a certain player has been authenticated.
// @param playerSlot: The index of the player's slot.
// @return true if the player is authenticated, false otherwise.
func IsClientAuthorized(playerSlot int32) bool {
	var __retVal bool
	__playerSlot := C.int32_t(playerSlot)
	__retVal = bool(C.IsClientAuthorized(__playerSlot))
	return __retVal
}

// IsClientConnected - Checks if a certain player is connected.
// @param playerSlot: The index of the player's slot.
// @return true if the player is connected, false otherwise.
func IsClientConnected(playerSlot int32) bool {
	var __retVal bool
	__playerSlot := C.int32_t(playerSlot)
	__retVal = bool(C.IsClientConnected(__playerSlot))
	return __retVal
}

// IsClientInGame - Checks if a certain player has entered the game.
// @param playerSlot: The index of the player's slot.
// @return true if the player is in the game, false otherwise.
func IsClientInGame(playerSlot int32) bool {
	var __retVal bool
	__playerSlot := C.int32_t(playerSlot)
	__retVal = bool(C.IsClientInGame(__playerSlot))
	return __retVal
}

// IsClientSourceTV - Checks if a certain player is the SourceTV bot.
// @param playerSlot: The index of the player's slot.
// @return true if the client is the SourceTV bot, false otherwise.
func IsClientSourceTV(playerSlot int32) bool {
	var __retVal bool
	__playerSlot := C.int32_t(playerSlot)
	__retVal = bool(C.IsClientSourceTV(__playerSlot))
	return __retVal
}

// IsClientAlive - Checks if the client is alive or dead.
// @param playerSlot: The index of the player's slot.
// @return true if the client is alive, false if dead.
func IsClientAlive(playerSlot int32) bool {
	var __retVal bool
	__playerSlot := C.int32_t(playerSlot)
	__retVal = bool(C.IsClientAlive(__playerSlot))
	return __retVal
}

// IsFakeClient - Checks if a certain player is a fake client.
// @param playerSlot: The index of the player's slot.
// @return true if the client is a fake client, false otherwise.
func IsFakeClient(playerSlot int32) bool {
	var __retVal bool
	__playerSlot := C.int32_t(playerSlot)
	__retVal = bool(C.IsFakeClient(__playerSlot))
	return __retVal
}

// GetClientTeam - Retrieves a client's team index.
// @param playerSlot: The index of the player's slot.
// @return int The team index of the client.
func GetClientTeam(playerSlot int32) int32 {
	var __retVal int32
	__playerSlot := C.int32_t(playerSlot)
	__retVal = int32(C.GetClientTeam(__playerSlot))
	return __retVal
}

// GetClientHealth - Returns the client's health.
// @param playerSlot: The index of the player's slot.
// @return int The health value of the client.
func GetClientHealth(playerSlot int32) int32 {
	var __retVal int32
	__playerSlot := C.int32_t(playerSlot)
	__retVal = int32(C.GetClientHealth(__playerSlot))
	return __retVal
}

// GetClientArmor - Returns the client's armor value.
// @param playerSlot: The index of the player's slot.
// @return int The armor value of the client.
func GetClientArmor(playerSlot int32) int32 {
	var __retVal int32
	__playerSlot := C.int32_t(playerSlot)
	__retVal = int32(C.GetClientArmor(__playerSlot))
	return __retVal
}

// GetClientAbsOrigin - Retrieves the client's origin vector.
// @param playerSlot: The index of the player's slot.
// @return A Vector where the client's origin will be stored.
func GetClientAbsOrigin(playerSlot int32) plugify.Vector3 {
	var __retVal plugify.Vector3
	__playerSlot := C.int32_t(playerSlot)
	__native := C.GetClientAbsOrigin(__playerSlot)
	__retVal = *(*plugify.Vector3)(unsafe.Pointer(&__native))
	return __retVal
}

// GetClientAbsAngles - Retrieves the client's position angle.
// @param playerSlot: The index of the player's slot.
// @return A QAngle where the client's position angle will be stored.
func GetClientAbsAngles(playerSlot int32) plugify.Vector3 {
	var __retVal plugify.Vector3
	__playerSlot := C.int32_t(playerSlot)
	__native := C.GetClientAbsAngles(__playerSlot)
	__retVal = *(*plugify.Vector3)(unsafe.Pointer(&__native))
	return __retVal
}

// GetClientEyeAngles - Retrieves the client's eye angle.
// @param playerSlot: The index of the player's slot.
// @return A QAngle where the client's eye angle will be stored.
func GetClientEyeAngles(playerSlot int32) plugify.Vector3 {
	var __retVal plugify.Vector3
	__playerSlot := C.int32_t(playerSlot)
	__native := C.GetClientEyeAngles(__playerSlot)
	__retVal = *(*plugify.Vector3)(unsafe.Pointer(&__native))
	return __retVal
}

// ProcessTargetString - Processes the target string to determine if one user can target another.
// @param caller: The index of the player's slot making the target request.
// @param target: The target string specifying the player or players to be targeted.
// @return A vector where the result of the targeting operation will be stored.
func ProcessTargetString(caller int32, target string) []int32 {
	var __retVal []int32
	var __retVal_native plugify.PlgVector
	__caller := C.int32_t(caller)
	__target := plugify.ConstructString(target)
	plugify.Block{
		Try: func() {
			__native := C.ProcessTargetString(__caller, (*C.String)(unsafe.Pointer(&__target)))
			__retVal_native = *(*plugify.PlgVector)(unsafe.Pointer(&__native))
			// Unmarshal - Convert native data to managed data.
			__retVal = plugify.GetVectorDataInt32(&__retVal_native)
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyVectorInt32(&__retVal_native)
			plugify.DestroyString(&__target)
		},
	}.Do()
	return __retVal
}

// ChangeClientTeam - Changes a client's team.
// @param playerSlot: The index of the player's slot.
// @param team: The team index to assign the client to.
func ChangeClientTeam(playerSlot int32, team CSTeam) {
	__playerSlot := C.int32_t(playerSlot)
	__team := C.int32_t(team)
	C.ChangeClientTeam(__playerSlot, __team)
}

// SwitchClientTeam - Switches the player's team.
// @param playerSlot: The index of the player's slot.
// @param team: The team index to switch the client to.
func SwitchClientTeam(playerSlot int32, team CSTeam) {
	__playerSlot := C.int32_t(playerSlot)
	__team := C.int32_t(team)
	C.SwitchClientTeam(__playerSlot, __team)
}

// RespawnClient - Respawns a player.
// @param playerSlot: The index of the player's slot to respawn.
func RespawnClient(playerSlot int32) {
	__playerSlot := C.int32_t(playerSlot)
	C.RespawnClient(__playerSlot)
}

// ForcePlayerSuicide - Forces a player to commit suicide.
// @param playerSlot: The index of the player's slot.
// @param explode: If true, the client will explode upon death.
// @param force: If true, the suicide will be forced.
func ForcePlayerSuicide(playerSlot int32, explode bool, force bool) {
	__playerSlot := C.int32_t(playerSlot)
	__explode := C.bool(explode)
	__force := C.bool(force)
	C.ForcePlayerSuicide(__playerSlot, __explode, __force)
}

// KickClient - Disconnects a client from the server as soon as the next frame starts.
// @param playerSlot: The index of the player's slot to be kicked.
func KickClient(playerSlot int32) {
	__playerSlot := C.int32_t(playerSlot)
	C.KickClient(__playerSlot)
}

// BanClient - Bans a client for a specified duration.
// @param playerSlot: The index of the player's slot to be banned.
// @param duration: Duration of the ban in seconds.
// @param kick: If true, the client will be kicked immediately after being banned.
func BanClient(playerSlot int32, duration float32, kick bool) {
	__playerSlot := C.int32_t(playerSlot)
	__duration := C.float(duration)
	__kick := C.bool(kick)
	C.BanClient(__playerSlot, __duration, __kick)
}

// BanIdentity - Bans an identity (either an IP address or a Steam authentication string).
// @param steamId: The Steam ID to ban.
// @param duration: Duration of the ban in seconds.
// @param kick: If true, the client will be kicked immediately after being banned.
func BanIdentity(steamId uint64, duration float32, kick bool) {
	__steamId := C.uint64_t(steamId)
	__duration := C.float(duration)
	__kick := C.bool(kick)
	C.BanIdentity(__steamId, __duration, __kick)
}

// GetClientActiveWeapon - Retrieves the handle of the client's currently active weapon.
// @param playerSlot: The index of the client.
// @return The entity handle of the active weapon, or INVALID_EHANDLE_INDEX if the client is invalid or has no active weapon.
func GetClientActiveWeapon(playerSlot int32) int32 {
	var __retVal int32
	__playerSlot := C.int32_t(playerSlot)
	__retVal = int32(C.GetClientActiveWeapon(__playerSlot))
	return __retVal
}

// GetClientWeapons - Retrieves a list of weapon handles owned by the client.
// @param playerSlot: The index of the client.
// @return A vector of entity handles for the client's weapons, or an empty vector if the client is invalid or has no weapons.
func GetClientWeapons(playerSlot int32) []int32 {
	var __retVal []int32
	var __retVal_native plugify.PlgVector
	__playerSlot := C.int32_t(playerSlot)
	plugify.Block{
		Try: func() {
			__native := C.GetClientWeapons(__playerSlot)
			__retVal_native = *(*plugify.PlgVector)(unsafe.Pointer(&__native))
			// Unmarshal - Convert native data to managed data.
			__retVal = plugify.GetVectorDataInt32(&__retVal_native)
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyVectorInt32(&__retVal_native)
		},
	}.Do()
	return __retVal
}

// DropWeapon - Forces a player to drop their weapon.
// @param playerSlot: The index of the player's slot.
// @param weaponHandle: The handle of weapon to drop.
// @param target: Target direction.
// @param velocity: Velocity to toss weapon or zero to just drop weapon.
func DropWeapon(playerSlot int32, weaponHandle int32, target plugify.Vector3, velocity plugify.Vector3) {
	__playerSlot := C.int32_t(playerSlot)
	__weaponHandle := C.int32_t(weaponHandle)
	__target := *(*C.Vector3)(unsafe.Pointer(&target))
	__velocity := *(*C.Vector3)(unsafe.Pointer(&velocity))
	C.DropWeapon(__playerSlot, __weaponHandle, &__target, &__velocity)
}

// StripWeapons - Removes all weapons from a client, with an option to remove the suit as well.
// @param playerSlot: The index of the client.
// @param removeSuit: A boolean indicating whether to also remove the client's suit.
func StripWeapons(playerSlot int32, removeSuit bool) {
	__playerSlot := C.int32_t(playerSlot)
	__removeSuit := C.bool(removeSuit)
	C.StripWeapons(__playerSlot, __removeSuit)
}

// BumpWeapon - Bumps a player's weapon.
// @param playerSlot: The index of the client.
// @param weaponHandle: The handle of weapon to bump.
func BumpWeapon(playerSlot int32, weaponHandle int32) {
	__playerSlot := C.int32_t(playerSlot)
	__weaponHandle := C.int32_t(weaponHandle)
	C.BumpWeapon(__playerSlot, __weaponHandle)
}

// SwitchWeapon - Switches a player's weapon.
// @param playerSlot: The index of the client.
// @param weaponHandle: The handle of weapon to switch.
func SwitchWeapon(playerSlot int32, weaponHandle int32) {
	__playerSlot := C.int32_t(playerSlot)
	__weaponHandle := C.int32_t(weaponHandle)
	C.SwitchWeapon(__playerSlot, __weaponHandle)
}

// RemoveWeapon - Removes a player's weapon.
// @param playerSlot: The index of the client.
// @param weaponHandle: The handle of weapon to remove.
func RemoveWeapon(playerSlot int32, weaponHandle int32) {
	__playerSlot := C.int32_t(playerSlot)
	__weaponHandle := C.int32_t(weaponHandle)
	C.RemoveWeapon(__playerSlot, __weaponHandle)
}

// GiveNamedItem - Gives a named item (e.g., weapon) to a client.
// @param playerSlot: The index of the client.
// @param itemName: The name of the item to give.
// @return The entity handle of the created item, or INVALID_EHANDLE_INDEX if the client or item is invalid.
func GiveNamedItem(playerSlot int32, itemName string) int32 {
	var __retVal int32
	__playerSlot := C.int32_t(playerSlot)
	__itemName := plugify.ConstructString(itemName)
	plugify.Block{
		Try: func() {
			__retVal = int32(C.GiveNamedItem(__playerSlot, (*C.String)(unsafe.Pointer(&__itemName))))
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__itemName)
		},
	}.Do()
	return __retVal
}

// GetClientButtons - Retrieves the state of a specific button for a client.
// @param playerSlot: The index of the client.
// @param buttonIndex: The index of the button (0-2).
// @return The state of the specified button, or 0 if the client or button index is invalid.
func GetClientButtons(playerSlot int32, buttonIndex int32) uint64 {
	var __retVal uint64
	__playerSlot := C.int32_t(playerSlot)
	__buttonIndex := C.int32_t(buttonIndex)
	__retVal = uint64(C.GetClientButtons(__playerSlot, __buttonIndex))
	return __retVal
}

// GetClientMoney - Retrieves the amount of money a client has.
// @param playerSlot: The index of the client.
// @return The amount of money the client has, or 0 if The player slot is invalid.
func GetClientMoney(playerSlot int32) int32 {
	var __retVal int32
	__playerSlot := C.int32_t(playerSlot)
	__retVal = int32(C.GetClientMoney(__playerSlot))
	return __retVal
}

// SetClientMoney - Sets the amount of money for a client.
// @param playerSlot: The index of the client.
// @param money: The amount of money to set.
func SetClientMoney(playerSlot int32, money int32) {
	__playerSlot := C.int32_t(playerSlot)
	__money := C.int32_t(money)
	C.SetClientMoney(__playerSlot, __money)
}

// GetClientKills - Retrieves the number of kills for a client.
// @param playerSlot: The index of the client.
// @return The number of kills the client has, or 0 if The player slot is invalid.
func GetClientKills(playerSlot int32) int32 {
	var __retVal int32
	__playerSlot := C.int32_t(playerSlot)
	__retVal = int32(C.GetClientKills(__playerSlot))
	return __retVal
}

// SetClientKills - Sets the number of kills for a client.
// @param playerSlot: The index of the client.
// @param kills: The number of kills to set.
func SetClientKills(playerSlot int32, kills int32) {
	__playerSlot := C.int32_t(playerSlot)
	__kills := C.int32_t(kills)
	C.SetClientKills(__playerSlot, __kills)
}

// GetClientDeaths - Retrieves the number of deaths for a client.
// @param playerSlot: The index of the client.
// @return The number of deaths the client has, or 0 if The player slot is invalid.
func GetClientDeaths(playerSlot int32) int32 {
	var __retVal int32
	__playerSlot := C.int32_t(playerSlot)
	__retVal = int32(C.GetClientDeaths(__playerSlot))
	return __retVal
}

// SetClientDeaths - Sets the number of deaths for a client.
// @param playerSlot: The index of the client.
// @param deaths: The number of deaths to set.
func SetClientDeaths(playerSlot int32, deaths int32) {
	__playerSlot := C.int32_t(playerSlot)
	__deaths := C.int32_t(deaths)
	C.SetClientDeaths(__playerSlot, __deaths)
}

// GetClientAssists - Retrieves the number of assists for a client.
// @param playerSlot: The index of the client.
// @return The number of assists the client has, or 0 if The player slot is invalid.
func GetClientAssists(playerSlot int32) int32 {
	var __retVal int32
	__playerSlot := C.int32_t(playerSlot)
	__retVal = int32(C.GetClientAssists(__playerSlot))
	return __retVal
}

// SetClientAssists - Sets the number of assists for a client.
// @param playerSlot: The index of the client.
// @param assists: The number of assists to set.
func SetClientAssists(playerSlot int32, assists int32) {
	__playerSlot := C.int32_t(playerSlot)
	__assists := C.int32_t(assists)
	C.SetClientAssists(__playerSlot, __assists)
}

// GetClientDamage - Retrieves the total damage dealt by a client.
// @param playerSlot: The index of the client.
// @return The total damage dealt by the client, or 0 if The player slot is invalid.
func GetClientDamage(playerSlot int32) int32 {
	var __retVal int32
	__playerSlot := C.int32_t(playerSlot)
	__retVal = int32(C.GetClientDamage(__playerSlot))
	return __retVal
}

// SetClientDamage - Sets the total damage dealt by a client.
// @param playerSlot: The index of the client.
// @param damage: The amount of damage to set.
func SetClientDamage(playerSlot int32, damage int32) {
	__playerSlot := C.int32_t(playerSlot)
	__damage := C.int32_t(damage)
	C.SetClientDamage(__playerSlot, __damage)
}

// AddAdminCommand - Creates a console command as an administrative command.
// @param name: The name of the console command.
// @param adminFlags: The admin flags that indicate which admin level can use this command.
// @param description: A brief description of what the command does.
// @param flags: Command flags that define the behavior of the command.
// @param callback: A callback function that is invoked when the command is executed.
// @param type: Whether the hook was in post mode (after processing) or pre mode (before processing).
// @return true if the command was successfully created; otherwise, false.
func AddAdminCommand(name string, adminFlags int64, description string, flags ConVarFlag, callback CommandCallback, type_ HookMode) bool {
	var __retVal bool
	__name := plugify.ConstructString(name)
	__adminFlags := C.int64_t(adminFlags)
	__description := plugify.ConstructString(description)
	__flags := C.int64_t(flags)
	__callback := plugify.GetFunctionPointerForDelegate(callback)
	__type_ := C.uint8_t(type_)
	plugify.Block{
		Try: func() {
			__retVal = bool(C.AddAdminCommand((*C.String)(unsafe.Pointer(&__name)), __adminFlags, (*C.String)(unsafe.Pointer(&__description)), __flags, __callback, __type_))
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__name)
			plugify.DestroyString(&__description)
		},
	}.Do()
	return __retVal
}

// AddConsoleCommand - Creates a console command or hooks an already existing one.
// @param name: The name of the console command.
// @param description: A brief description of what the command does.
// @param flags: Command flags that define the behavior of the command.
// @param callback: A callback function that is invoked when the command is executed.
// @param type: Whether the hook was in post mode (after processing) or pre mode (before processing).
// @return true if the command was successfully created; otherwise, false.
func AddConsoleCommand(name string, description string, flags ConVarFlag, callback CommandCallback, type_ HookMode) bool {
	var __retVal bool
	__name := plugify.ConstructString(name)
	__description := plugify.ConstructString(description)
	__flags := C.int64_t(flags)
	__callback := plugify.GetFunctionPointerForDelegate(callback)
	__type_ := C.uint8_t(type_)
	plugify.Block{
		Try: func() {
			__retVal = bool(C.AddConsoleCommand((*C.String)(unsafe.Pointer(&__name)), (*C.String)(unsafe.Pointer(&__description)), __flags, __callback, __type_))
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__name)
			plugify.DestroyString(&__description)
		},
	}.Do()
	return __retVal
}

// RemoveCommand - Removes a console command from the system.
// @param name: The name of the command to be removed.
// @param callback: The callback function associated with the command to be removed.
// @return true if the command was successfully removed; otherwise, false.
func RemoveCommand(name string, callback CommandCallback) bool {
	var __retVal bool
	__name := plugify.ConstructString(name)
	__callback := plugify.GetFunctionPointerForDelegate(callback)
	plugify.Block{
		Try: func() {
			__retVal = bool(C.RemoveCommand((*C.String)(unsafe.Pointer(&__name)), __callback))
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__name)
		},
	}.Do()
	return __retVal
}

// AddCommandListener - Adds a callback that will fire when a command is sent to the server.
// @param name: The name of the command.
// @param callback: The callback function that will be invoked when the command is executed.
// @param type: Whether the hook was in post mode (after processing) or pre mode (before processing).
// @return Returns true if the callback was successfully added, false otherwise.
func AddCommandListener(name string, callback CommandCallback, type_ HookMode) bool {
	var __retVal bool
	__name := plugify.ConstructString(name)
	__callback := plugify.GetFunctionPointerForDelegate(callback)
	__type_ := C.uint8_t(type_)
	plugify.Block{
		Try: func() {
			__retVal = bool(C.AddCommandListener((*C.String)(unsafe.Pointer(&__name)), __callback, __type_))
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__name)
		},
	}.Do()
	return __retVal
}

// RemoveCommandListener - Removes a callback that fires when a command is sent to the server.
// @param name: The name of the command.
// @param callback: The callback function to be removed.
// @param type: Whether the hook was in post mode (after processing) or pre mode (before processing).
// @return Returns true if the callback was successfully removed, false otherwise.
func RemoveCommandListener(name string, callback CommandCallback, type_ HookMode) bool {
	var __retVal bool
	__name := plugify.ConstructString(name)
	__callback := plugify.GetFunctionPointerForDelegate(callback)
	__type_ := C.uint8_t(type_)
	plugify.Block{
		Try: func() {
			__retVal = bool(C.RemoveCommandListener((*C.String)(unsafe.Pointer(&__name)), __callback, __type_))
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__name)
		},
	}.Do()
	return __retVal
}

// ServerCommand - Executes a server command as if it were run on the server console or through RCON.
// @param command: The command to execute on the server.
func ServerCommand(command string) {
	__command := plugify.ConstructString(command)
	plugify.Block{
		Try: func() {
			C.ServerCommand((*C.String)(unsafe.Pointer(&__command)))
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__command)
		},
	}.Do()
}

// ServerCommandEx - Executes a server command as if it were on the server console (or RCON) and stores the printed text into buffer.
// @param command: The command to execute on the server.
// @return String to store command result into.
func ServerCommandEx(command string) string {
	var __retVal string
	var __retVal_native plugify.PlgString
	__command := plugify.ConstructString(command)
	plugify.Block{
		Try: func() {
			__native := C.ServerCommandEx((*C.String)(unsafe.Pointer(&__command)))
			__retVal_native = *(*plugify.PlgString)(unsafe.Pointer(&__native))
			// Unmarshal - Convert native data to managed data.
			__retVal = plugify.GetStringData(&__retVal_native)
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__retVal_native)
			plugify.DestroyString(&__command)
		},
	}.Do()
	return __retVal
}

// ClientCommand - Executes a client command.
// @param playerSlot: The index of the client executing the command.
// @param command: The command to execute on the client.
func ClientCommand(playerSlot int32, command string) {
	__playerSlot := C.int32_t(playerSlot)
	__command := plugify.ConstructString(command)
	plugify.Block{
		Try: func() {
			C.ClientCommand(__playerSlot, (*C.String)(unsafe.Pointer(&__command)))
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__command)
		},
	}.Do()
}

// FakeClientCommand - Executes a client command on the server without network communication.
// @param playerSlot: The index of the client.
// @param command: The command to be executed by the client.
func FakeClientCommand(playerSlot int32, command string) {
	__playerSlot := C.int32_t(playerSlot)
	__command := plugify.ConstructString(command)
	plugify.Block{
		Try: func() {
			C.FakeClientCommand(__playerSlot, (*C.String)(unsafe.Pointer(&__command)))
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__command)
		},
	}.Do()
}

// PrintToServer - Sends a message to the server console.
// @param msg: The message to be sent to the server console.
func PrintToServer(msg string) {
	__msg := plugify.ConstructString(msg)
	plugify.Block{
		Try: func() {
			C.PrintToServer((*C.String)(unsafe.Pointer(&__msg)))
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__msg)
		},
	}.Do()
}

// PrintToConsole - Sends a message to a client's console.
// @param playerSlot: The index of the player's slot to whom the message will be sent.
// @param message: The message to be sent to the client's console.
func PrintToConsole(playerSlot int32, message string) {
	__playerSlot := C.int32_t(playerSlot)
	__message := plugify.ConstructString(message)
	plugify.Block{
		Try: func() {
			C.PrintToConsole(__playerSlot, (*C.String)(unsafe.Pointer(&__message)))
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__message)
		},
	}.Do()
}

// PrintToChat - Prints a message to a specific client in the chat area.
// @param playerSlot: The index of the player's slot to whom the message will be sent.
// @param message: The message to be printed in the chat area.
func PrintToChat(playerSlot int32, message string) {
	__playerSlot := C.int32_t(playerSlot)
	__message := plugify.ConstructString(message)
	plugify.Block{
		Try: func() {
			C.PrintToChat(__playerSlot, (*C.String)(unsafe.Pointer(&__message)))
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__message)
		},
	}.Do()
}

// PrintCenterText - Prints a message to a specific client in the center of the screen.
// @param playerSlot: The index of the player's slot to whom the message will be sent.
// @param message: The message to be printed in the center of the screen.
func PrintCenterText(playerSlot int32, message string) {
	__playerSlot := C.int32_t(playerSlot)
	__message := plugify.ConstructString(message)
	plugify.Block{
		Try: func() {
			C.PrintCenterText(__playerSlot, (*C.String)(unsafe.Pointer(&__message)))
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__message)
		},
	}.Do()
}

// PrintAlertText - Prints a message to a specific client with an alert box.
// @param playerSlot: The index of the player's slot to whom the message will be sent.
// @param message: The message to be printed in the alert box.
func PrintAlertText(playerSlot int32, message string) {
	__playerSlot := C.int32_t(playerSlot)
	__message := plugify.ConstructString(message)
	plugify.Block{
		Try: func() {
			C.PrintAlertText(__playerSlot, (*C.String)(unsafe.Pointer(&__message)))
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__message)
		},
	}.Do()
}

// PrintCentreHtml - Prints a html message to a specific client in the center of the screen.
// @param playerSlot: The index of the player's slot to whom the message will be sent.
// @param message: The HTML-formatted message to be printed.
func PrintCentreHtml(playerSlot int32, message string) {
	__playerSlot := C.int32_t(playerSlot)
	__message := plugify.ConstructString(message)
	plugify.Block{
		Try: func() {
			C.PrintCentreHtml(__playerSlot, (*C.String)(unsafe.Pointer(&__message)))
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__message)
		},
	}.Do()
}

// PrintToConsoleAll - Sends a message to every client's console.
// @param message: The message to be sent to all clients' consoles.
func PrintToConsoleAll(message string) {
	__message := plugify.ConstructString(message)
	plugify.Block{
		Try: func() {
			C.PrintToConsoleAll((*C.String)(unsafe.Pointer(&__message)))
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__message)
		},
	}.Do()
}

// PrintToChatAll - Prints a message to all clients in the chat area.
// @param message: The message to be printed in the chat area for all clients.
func PrintToChatAll(message string) {
	__message := plugify.ConstructString(message)
	plugify.Block{
		Try: func() {
			C.PrintToChatAll((*C.String)(unsafe.Pointer(&__message)))
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__message)
		},
	}.Do()
}

// PrintCenterTextAll - Prints a message to all clients in the center of the screen.
// @param message: The message to be printed in the center of the screen for all clients.
func PrintCenterTextAll(message string) {
	__message := plugify.ConstructString(message)
	plugify.Block{
		Try: func() {
			C.PrintCenterTextAll((*C.String)(unsafe.Pointer(&__message)))
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__message)
		},
	}.Do()
}

// PrintAlertTextAll - Prints a message to all clients with an alert box.
// @param message: The message to be printed in an alert box for all clients.
func PrintAlertTextAll(message string) {
	__message := plugify.ConstructString(message)
	plugify.Block{
		Try: func() {
			C.PrintAlertTextAll((*C.String)(unsafe.Pointer(&__message)))
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__message)
		},
	}.Do()
}

// PrintCentreHtmlAll - Prints a html message to all clients in the center of the screen.
// @param message: The HTML-formatted message to be printed in the center of the screen for all clients.
func PrintCentreHtmlAll(message string) {
	__message := plugify.ConstructString(message)
	plugify.Block{
		Try: func() {
			C.PrintCentreHtmlAll((*C.String)(unsafe.Pointer(&__message)))
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__message)
		},
	}.Do()
}

// PrintToChatColored - Prints a colored message to a specific client in the chat area.
// @param playerSlot: The index of the player's slot to whom the message will be sent.
// @param message: The message to be printed in the chat area with color.
func PrintToChatColored(playerSlot int32, message string) {
	__playerSlot := C.int32_t(playerSlot)
	__message := plugify.ConstructString(message)
	plugify.Block{
		Try: func() {
			C.PrintToChatColored(__playerSlot, (*C.String)(unsafe.Pointer(&__message)))
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__message)
		},
	}.Do()
}

// PrintToChatColoredAll - Prints a colored message to all clients in the chat area.
// @param message: The colored message to be printed in the chat area for all clients.
func PrintToChatColoredAll(message string) {
	__message := plugify.ConstructString(message)
	plugify.Block{
		Try: func() {
			C.PrintToChatColoredAll((*C.String)(unsafe.Pointer(&__message)))
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__message)
		},
	}.Do()
}

// CreateConVar - Creates a new console variable.
// @param name: The name of the console variable.
// @param defaultValue: The default value of the console variable.
// @param description: A description of the console variable's purpose.
// @param flags: Additional flags for the console variable.
// @return A handle to the created console variable.
func CreateConVar(name string, defaultValue string, description string, flags ConVarFlag) uint64 {
	var __retVal uint64
	__name := plugify.ConstructString(name)
	__defaultValue := plugify.ConstructString(defaultValue)
	__description := plugify.ConstructString(description)
	__flags := C.int64_t(flags)
	plugify.Block{
		Try: func() {
			__retVal = uint64(C.CreateConVar((*C.String)(unsafe.Pointer(&__name)), (*C.String)(unsafe.Pointer(&__defaultValue)), (*C.String)(unsafe.Pointer(&__description)), __flags))
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__name)
			plugify.DestroyString(&__defaultValue)
			plugify.DestroyString(&__description)
		},
	}.Do()
	return __retVal
}

// CreateConVarBool - Creates a new boolean console variable.
// @param name: The name of the console variable.
// @param defaultValue: The default value for the console variable.
// @param description: A brief description of the console variable.
// @param flags: Flags that define the behavior of the console variable.
// @param hasMin: Indicates if a minimum value is provided.
// @param min: The minimum value if hasMin is true.
// @param hasMax: Indicates if a maximum value is provided.
// @param max: The maximum value if hasMax is true.
// @return A handle to the created console variable data.
func CreateConVarBool(name string, defaultValue bool, description string, flags ConVarFlag, hasMin bool, min bool, hasMax bool, max bool) uint64 {
	var __retVal uint64
	__name := plugify.ConstructString(name)
	__defaultValue := C.bool(defaultValue)
	__description := plugify.ConstructString(description)
	__flags := C.int64_t(flags)
	__hasMin := C.bool(hasMin)
	__min := C.bool(min)
	__hasMax := C.bool(hasMax)
	__max := C.bool(max)
	plugify.Block{
		Try: func() {
			__retVal = uint64(C.CreateConVarBool((*C.String)(unsafe.Pointer(&__name)), __defaultValue, (*C.String)(unsafe.Pointer(&__description)), __flags, __hasMin, __min, __hasMax, __max))
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__name)
			plugify.DestroyString(&__description)
		},
	}.Do()
	return __retVal
}

// CreateConVarInt16 - Creates a new 16-bit signed integer console variable.
// @param name: The name of the console variable.
// @param defaultValue: The default value for the console variable.
// @param description: A brief description of the console variable.
// @param flags: Flags that define the behavior of the console variable.
// @param hasMin: Indicates if a minimum value is provided.
// @param min: The minimum value if hasMin is true.
// @param hasMax: Indicates if a maximum value is provided.
// @param max: The maximum value if hasMax is true.
// @return A handle to the created console variable data.
func CreateConVarInt16(name string, defaultValue int16, description string, flags ConVarFlag, hasMin bool, min int16, hasMax bool, max int16) uint64 {
	var __retVal uint64
	__name := plugify.ConstructString(name)
	__defaultValue := C.int16_t(defaultValue)
	__description := plugify.ConstructString(description)
	__flags := C.int64_t(flags)
	__hasMin := C.bool(hasMin)
	__min := C.int16_t(min)
	__hasMax := C.bool(hasMax)
	__max := C.int16_t(max)
	plugify.Block{
		Try: func() {
			__retVal = uint64(C.CreateConVarInt16((*C.String)(unsafe.Pointer(&__name)), __defaultValue, (*C.String)(unsafe.Pointer(&__description)), __flags, __hasMin, __min, __hasMax, __max))
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__name)
			plugify.DestroyString(&__description)
		},
	}.Do()
	return __retVal
}

// CreateConVarUInt16 - Creates a new 16-bit unsigned integer console variable.
// @param name: The name of the console variable.
// @param defaultValue: The default value for the console variable.
// @param description: A brief description of the console variable.
// @param flags: Flags that define the behavior of the console variable.
// @param hasMin: Indicates if a minimum value is provided.
// @param min: The minimum value if hasMin is true.
// @param hasMax: Indicates if a maximum value is provided.
// @param max: The maximum value if hasMax is true.
// @return A handle to the created console variable data.
func CreateConVarUInt16(name string, defaultValue uint16, description string, flags ConVarFlag, hasMin bool, min uint16, hasMax bool, max uint16) uint64 {
	var __retVal uint64
	__name := plugify.ConstructString(name)
	__defaultValue := C.uint16_t(defaultValue)
	__description := plugify.ConstructString(description)
	__flags := C.int64_t(flags)
	__hasMin := C.bool(hasMin)
	__min := C.uint16_t(min)
	__hasMax := C.bool(hasMax)
	__max := C.uint16_t(max)
	plugify.Block{
		Try: func() {
			__retVal = uint64(C.CreateConVarUInt16((*C.String)(unsafe.Pointer(&__name)), __defaultValue, (*C.String)(unsafe.Pointer(&__description)), __flags, __hasMin, __min, __hasMax, __max))
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__name)
			plugify.DestroyString(&__description)
		},
	}.Do()
	return __retVal
}

// CreateConVarInt32 - Creates a new 32-bit signed integer console variable.
// @param name: The name of the console variable.
// @param defaultValue: The default value for the console variable.
// @param description: A brief description of the console variable.
// @param flags: Flags that define the behavior of the console variable.
// @param hasMin: Indicates if a minimum value is provided.
// @param min: The minimum value if hasMin is true.
// @param hasMax: Indicates if a maximum value is provided.
// @param max: The maximum value if hasMax is true.
// @return A handle to the created console variable data.
func CreateConVarInt32(name string, defaultValue int32, description string, flags ConVarFlag, hasMin bool, min int32, hasMax bool, max int32) uint64 {
	var __retVal uint64
	__name := plugify.ConstructString(name)
	__defaultValue := C.int32_t(defaultValue)
	__description := plugify.ConstructString(description)
	__flags := C.int64_t(flags)
	__hasMin := C.bool(hasMin)
	__min := C.int32_t(min)
	__hasMax := C.bool(hasMax)
	__max := C.int32_t(max)
	plugify.Block{
		Try: func() {
			__retVal = uint64(C.CreateConVarInt32((*C.String)(unsafe.Pointer(&__name)), __defaultValue, (*C.String)(unsafe.Pointer(&__description)), __flags, __hasMin, __min, __hasMax, __max))
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__name)
			plugify.DestroyString(&__description)
		},
	}.Do()
	return __retVal
}

// CreateConVarUInt32 - Creates a new 32-bit unsigned integer console variable.
// @param name: The name of the console variable.
// @param defaultValue: The default value for the console variable.
// @param description: A brief description of the console variable.
// @param flags: Flags that define the behavior of the console variable.
// @param hasMin: Indicates if a minimum value is provided.
// @param min: The minimum value if hasMin is true.
// @param hasMax: Indicates if a maximum value is provided.
// @param max: The maximum value if hasMax is true.
// @return A handle to the created console variable data.
func CreateConVarUInt32(name string, defaultValue uint32, description string, flags ConVarFlag, hasMin bool, min uint32, hasMax bool, max uint32) uint64 {
	var __retVal uint64
	__name := plugify.ConstructString(name)
	__defaultValue := C.uint32_t(defaultValue)
	__description := plugify.ConstructString(description)
	__flags := C.int64_t(flags)
	__hasMin := C.bool(hasMin)
	__min := C.uint32_t(min)
	__hasMax := C.bool(hasMax)
	__max := C.uint32_t(max)
	plugify.Block{
		Try: func() {
			__retVal = uint64(C.CreateConVarUInt32((*C.String)(unsafe.Pointer(&__name)), __defaultValue, (*C.String)(unsafe.Pointer(&__description)), __flags, __hasMin, __min, __hasMax, __max))
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__name)
			plugify.DestroyString(&__description)
		},
	}.Do()
	return __retVal
}

// CreateConVarInt64 - Creates a new 64-bit signed integer console variable.
// @param name: The name of the console variable.
// @param defaultValue: The default value for the console variable.
// @param description: A brief description of the console variable.
// @param flags: Flags that define the behavior of the console variable.
// @param hasMin: Indicates if a minimum value is provided.
// @param min: The minimum value if hasMin is true.
// @param hasMax: Indicates if a maximum value is provided.
// @param max: The maximum value if hasMax is true.
// @return A handle to the created console variable data.
func CreateConVarInt64(name string, defaultValue int64, description string, flags ConVarFlag, hasMin bool, min int64, hasMax bool, max int64) uint64 {
	var __retVal uint64
	__name := plugify.ConstructString(name)
	__defaultValue := C.int64_t(defaultValue)
	__description := plugify.ConstructString(description)
	__flags := C.int64_t(flags)
	__hasMin := C.bool(hasMin)
	__min := C.int64_t(min)
	__hasMax := C.bool(hasMax)
	__max := C.int64_t(max)
	plugify.Block{
		Try: func() {
			__retVal = uint64(C.CreateConVarInt64((*C.String)(unsafe.Pointer(&__name)), __defaultValue, (*C.String)(unsafe.Pointer(&__description)), __flags, __hasMin, __min, __hasMax, __max))
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__name)
			plugify.DestroyString(&__description)
		},
	}.Do()
	return __retVal
}

// CreateConVarUInt64 - Creates a new 64-bit unsigned integer console variable.
// @param name: The name of the console variable.
// @param defaultValue: The default value for the console variable.
// @param description: A brief description of the console variable.
// @param flags: Flags that define the behavior of the console variable.
// @param hasMin: Indicates if a minimum value is provided.
// @param min: The minimum value if hasMin is true.
// @param hasMax: Indicates if a maximum value is provided.
// @param max: The maximum value if hasMax is true.
// @return A handle to the created console variable data.
func CreateConVarUInt64(name string, defaultValue uint64, description string, flags ConVarFlag, hasMin bool, min uint64, hasMax bool, max uint64) uint64 {
	var __retVal uint64
	__name := plugify.ConstructString(name)
	__defaultValue := C.uint64_t(defaultValue)
	__description := plugify.ConstructString(description)
	__flags := C.int64_t(flags)
	__hasMin := C.bool(hasMin)
	__min := C.uint64_t(min)
	__hasMax := C.bool(hasMax)
	__max := C.uint64_t(max)
	plugify.Block{
		Try: func() {
			__retVal = uint64(C.CreateConVarUInt64((*C.String)(unsafe.Pointer(&__name)), __defaultValue, (*C.String)(unsafe.Pointer(&__description)), __flags, __hasMin, __min, __hasMax, __max))
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__name)
			plugify.DestroyString(&__description)
		},
	}.Do()
	return __retVal
}

// CreateConVarFloat - Creates a new floating-point console variable.
// @param name: The name of the console variable.
// @param defaultValue: The default value for the console variable.
// @param description: A brief description of the console variable.
// @param flags: Flags that define the behavior of the console variable.
// @param hasMin: Indicates if a minimum value is provided.
// @param min: The minimum value if hasMin is true.
// @param hasMax: Indicates if a maximum value is provided.
// @param max: The maximum value if hasMax is true.
// @return A handle to the created console variable data.
func CreateConVarFloat(name string, defaultValue float32, description string, flags ConVarFlag, hasMin bool, min float32, hasMax bool, max float32) uint64 {
	var __retVal uint64
	__name := plugify.ConstructString(name)
	__defaultValue := C.float(defaultValue)
	__description := plugify.ConstructString(description)
	__flags := C.int64_t(flags)
	__hasMin := C.bool(hasMin)
	__min := C.float(min)
	__hasMax := C.bool(hasMax)
	__max := C.float(max)
	plugify.Block{
		Try: func() {
			__retVal = uint64(C.CreateConVarFloat((*C.String)(unsafe.Pointer(&__name)), __defaultValue, (*C.String)(unsafe.Pointer(&__description)), __flags, __hasMin, __min, __hasMax, __max))
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__name)
			plugify.DestroyString(&__description)
		},
	}.Do()
	return __retVal
}

// CreateConVarDouble - Creates a new double-precision console variable.
// @param name: The name of the console variable.
// @param defaultValue: The default value for the console variable.
// @param description: A brief description of the console variable.
// @param flags: Flags that define the behavior of the console variable.
// @param hasMin: Indicates if a minimum value is provided.
// @param min: The minimum value if hasMin is true.
// @param hasMax: Indicates if a maximum value is provided.
// @param max: The maximum value if hasMax is true.
// @return A handle to the created console variable data.
func CreateConVarDouble(name string, defaultValue float64, description string, flags ConVarFlag, hasMin bool, min float64, hasMax bool, max float64) uint64 {
	var __retVal uint64
	__name := plugify.ConstructString(name)
	__defaultValue := C.double(defaultValue)
	__description := plugify.ConstructString(description)
	__flags := C.int64_t(flags)
	__hasMin := C.bool(hasMin)
	__min := C.double(min)
	__hasMax := C.bool(hasMax)
	__max := C.double(max)
	plugify.Block{
		Try: func() {
			__retVal = uint64(C.CreateConVarDouble((*C.String)(unsafe.Pointer(&__name)), __defaultValue, (*C.String)(unsafe.Pointer(&__description)), __flags, __hasMin, __min, __hasMax, __max))
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__name)
			plugify.DestroyString(&__description)
		},
	}.Do()
	return __retVal
}

// CreateConVarColor - Creates a new color console variable.
// @param name: The name of the console variable.
// @param defaultValue: The default color value for the console variable.
// @param description: A brief description of the console variable.
// @param flags: Flags that define the behavior of the console variable.
// @param hasMin: Indicates if a minimum value is provided.
// @param min: The minimum color value if hasMin is true.
// @param hasMax: Indicates if a maximum value is provided.
// @param max: The maximum color value if hasMax is true.
// @return A handle to the created console variable data.
func CreateConVarColor(name string, defaultValue int32, description string, flags ConVarFlag, hasMin bool, min int32, hasMax bool, max int32) uint64 {
	var __retVal uint64
	__name := plugify.ConstructString(name)
	__defaultValue := C.int32_t(defaultValue)
	__description := plugify.ConstructString(description)
	__flags := C.int64_t(flags)
	__hasMin := C.bool(hasMin)
	__min := C.int32_t(min)
	__hasMax := C.bool(hasMax)
	__max := C.int32_t(max)
	plugify.Block{
		Try: func() {
			__retVal = uint64(C.CreateConVarColor((*C.String)(unsafe.Pointer(&__name)), __defaultValue, (*C.String)(unsafe.Pointer(&__description)), __flags, __hasMin, __min, __hasMax, __max))
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__name)
			plugify.DestroyString(&__description)
		},
	}.Do()
	return __retVal
}

// CreateConVarVector2 - Creates a new 2D vector console variable.
// @param name: The name of the console variable.
// @param defaultValue: The default value for the console variable.
// @param description: A brief description of the console variable.
// @param flags: Flags that define the behavior of the console variable.
// @param hasMin: Indicates if a minimum value is provided.
// @param min: The minimum value if hasMin is true.
// @param hasMax: Indicates if a maximum value is provided.
// @param max: The maximum value if hasMax is true.
// @return A handle to the created console variable data.
func CreateConVarVector2(name string, defaultValue plugify.Vector2, description string, flags ConVarFlag, hasMin bool, min plugify.Vector2, hasMax bool, max plugify.Vector2) uint64 {
	var __retVal uint64
	__name := plugify.ConstructString(name)
	__defaultValue := *(*C.Vector2)(unsafe.Pointer(&defaultValue))
	__description := plugify.ConstructString(description)
	__flags := C.int64_t(flags)
	__hasMin := C.bool(hasMin)
	__min := *(*C.Vector2)(unsafe.Pointer(&min))
	__hasMax := C.bool(hasMax)
	__max := *(*C.Vector2)(unsafe.Pointer(&max))
	plugify.Block{
		Try: func() {
			__retVal = uint64(C.CreateConVarVector2((*C.String)(unsafe.Pointer(&__name)), &__defaultValue, (*C.String)(unsafe.Pointer(&__description)), __flags, __hasMin, &__min, __hasMax, &__max))
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__name)
			plugify.DestroyString(&__description)
		},
	}.Do()
	return __retVal
}

// CreateConVarVector3 - Creates a new 3D vector console variable.
// @param name: The name of the console variable.
// @param defaultValue: The default value for the console variable.
// @param description: A brief description of the console variable.
// @param flags: Flags that define the behavior of the console variable.
// @param hasMin: Indicates if a minimum value is provided.
// @param min: The minimum value if hasMin is true.
// @param hasMax: Indicates if a maximum value is provided.
// @param max: The maximum value if hasMax is true.
// @return A handle to the created console variable data.
func CreateConVarVector3(name string, defaultValue plugify.Vector3, description string, flags ConVarFlag, hasMin bool, min plugify.Vector3, hasMax bool, max plugify.Vector3) uint64 {
	var __retVal uint64
	__name := plugify.ConstructString(name)
	__defaultValue := *(*C.Vector3)(unsafe.Pointer(&defaultValue))
	__description := plugify.ConstructString(description)
	__flags := C.int64_t(flags)
	__hasMin := C.bool(hasMin)
	__min := *(*C.Vector3)(unsafe.Pointer(&min))
	__hasMax := C.bool(hasMax)
	__max := *(*C.Vector3)(unsafe.Pointer(&max))
	plugify.Block{
		Try: func() {
			__retVal = uint64(C.CreateConVarVector3((*C.String)(unsafe.Pointer(&__name)), &__defaultValue, (*C.String)(unsafe.Pointer(&__description)), __flags, __hasMin, &__min, __hasMax, &__max))
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__name)
			plugify.DestroyString(&__description)
		},
	}.Do()
	return __retVal
}

// CreateConVarVector4 - Creates a new 4D vector console variable.
// @param name: The name of the console variable.
// @param defaultValue: The default value for the console variable.
// @param description: A brief description of the console variable.
// @param flags: Flags that define the behavior of the console variable.
// @param hasMin: Indicates if a minimum value is provided.
// @param min: The minimum value if hasMin is true.
// @param hasMax: Indicates if a maximum value is provided.
// @param max: The maximum value if hasMax is true.
// @return A handle to the created console variable data.
func CreateConVarVector4(name string, defaultValue plugify.Vector4, description string, flags ConVarFlag, hasMin bool, min plugify.Vector4, hasMax bool, max plugify.Vector4) uint64 {
	var __retVal uint64
	__name := plugify.ConstructString(name)
	__defaultValue := *(*C.Vector4)(unsafe.Pointer(&defaultValue))
	__description := plugify.ConstructString(description)
	__flags := C.int64_t(flags)
	__hasMin := C.bool(hasMin)
	__min := *(*C.Vector4)(unsafe.Pointer(&min))
	__hasMax := C.bool(hasMax)
	__max := *(*C.Vector4)(unsafe.Pointer(&max))
	plugify.Block{
		Try: func() {
			__retVal = uint64(C.CreateConVarVector4((*C.String)(unsafe.Pointer(&__name)), &__defaultValue, (*C.String)(unsafe.Pointer(&__description)), __flags, __hasMin, &__min, __hasMax, &__max))
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__name)
			plugify.DestroyString(&__description)
		},
	}.Do()
	return __retVal
}

// CreateConVarQAngle - Creates a new quaternion angle console variable.
// @param name: The name of the console variable.
// @param defaultValue: The default value for the console variable.
// @param description: A brief description of the console variable.
// @param flags: Flags that define the behavior of the console variable.
// @param hasMin: Indicates if a minimum value is provided.
// @param min: The minimum value if hasMin is true.
// @param hasMax: Indicates if a maximum value is provided.
// @param max: The maximum value if hasMax is true.
// @return A handle to the created console variable data.
func CreateConVarQAngle(name string, defaultValue plugify.Vector3, description string, flags ConVarFlag, hasMin bool, min plugify.Vector3, hasMax bool, max plugify.Vector3) uint64 {
	var __retVal uint64
	__name := plugify.ConstructString(name)
	__defaultValue := *(*C.Vector3)(unsafe.Pointer(&defaultValue))
	__description := plugify.ConstructString(description)
	__flags := C.int64_t(flags)
	__hasMin := C.bool(hasMin)
	__min := *(*C.Vector3)(unsafe.Pointer(&min))
	__hasMax := C.bool(hasMax)
	__max := *(*C.Vector3)(unsafe.Pointer(&max))
	plugify.Block{
		Try: func() {
			__retVal = uint64(C.CreateConVarQAngle((*C.String)(unsafe.Pointer(&__name)), &__defaultValue, (*C.String)(unsafe.Pointer(&__description)), __flags, __hasMin, &__min, __hasMax, &__max))
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__name)
			plugify.DestroyString(&__description)
		},
	}.Do()
	return __retVal
}

// FindConVar - Searches for a console variable.
// @param name: The name of the console variable to search for.
// @return A handle to the console variable data if found; otherwise, nullptr.
func FindConVar(name string) uint64 {
	var __retVal uint64
	__name := plugify.ConstructString(name)
	plugify.Block{
		Try: func() {
			__retVal = uint64(C.FindConVar((*C.String)(unsafe.Pointer(&__name))))
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__name)
		},
	}.Do()
	return __retVal
}

// FindConVar2 - Searches for a console variable of a specific type.
// @param name: The name of the console variable to search for.
// @param type: The type of the console variable to search for.
// @return A handle to the console variable data if found; otherwise, nullptr.
func FindConVar2(name string, type_ ConVarType) uint64 {
	var __retVal uint64
	__name := plugify.ConstructString(name)
	__type_ := C.int16_t(type_)
	plugify.Block{
		Try: func() {
			__retVal = uint64(C.FindConVar2((*C.String)(unsafe.Pointer(&__name)), __type_))
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__name)
		},
	}.Do()
	return __retVal
}

// HookConVarChange - Creates a hook for when a console variable's value is changed.
// @param name: The name of the console variable to hook.
// @param callback: The callback function to be executed when the variable's value changes.
func HookConVarChange(name string, callback ChangeCallback) {
	__name := plugify.ConstructString(name)
	__callback := plugify.GetFunctionPointerForDelegate(callback)
	plugify.Block{
		Try: func() {
			C.HookConVarChange((*C.String)(unsafe.Pointer(&__name)), __callback)
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__name)
		},
	}.Do()
}

// UnhookConVarChange - Removes a hook for when a console variable's value is changed.
// @param name: The name of the console variable to unhook.
// @param callback: The callback function to be removed.
func UnhookConVarChange(name string, callback ChangeCallback) {
	__name := plugify.ConstructString(name)
	__callback := plugify.GetFunctionPointerForDelegate(callback)
	plugify.Block{
		Try: func() {
			C.UnhookConVarChange((*C.String)(unsafe.Pointer(&__name)), __callback)
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__name)
		},
	}.Do()
}

// IsConVarFlagSet - Checks if a specific flag is set for a console variable.
// @param conVarHandle: The handle to the console variable data.
// @param flag: The flag to check against the console variable.
// @return True if the flag is set; otherwise, false.
func IsConVarFlagSet(conVarHandle uint64, flag int64) bool {
	var __retVal bool
	__conVarHandle := C.uint64_t(conVarHandle)
	__flag := C.int64_t(flag)
	__retVal = bool(C.IsConVarFlagSet(__conVarHandle, __flag))
	return __retVal
}

// AddConVarFlags - Adds flags to a console variable.
// @param conVarHandle: The handle to the console variable data.
// @param flags: The flags to be added.
func AddConVarFlags(conVarHandle uint64, flags ConVarFlag) {
	__conVarHandle := C.uint64_t(conVarHandle)
	__flags := C.int64_t(flags)
	C.AddConVarFlags(__conVarHandle, __flags)
}

// RemoveConVarFlags - Removes flags from a console variable.
// @param conVarHandle: The handle to the console variable data.
// @param flags: The flags to be removed.
func RemoveConVarFlags(conVarHandle uint64, flags ConVarFlag) {
	__conVarHandle := C.uint64_t(conVarHandle)
	__flags := C.int64_t(flags)
	C.RemoveConVarFlags(__conVarHandle, __flags)
}

// GetConVarFlags - Retrieves the current flags of a console variable.
// @param conVarHandle: The handle to the console variable data.
// @return The current flags set on the console variable.
func GetConVarFlags(conVarHandle uint64) ConVarFlag {
	var __retVal ConVarFlag
	__conVarHandle := C.uint64_t(conVarHandle)
	__retVal = int64(C.GetConVarFlags(__conVarHandle))
	return __retVal
}

// GetConVarBounds - Gets the specified bound (max or min) of a console variable and stores it in the output string.
// @param conVarHandle: The handle to the console variable data.
// @param max: Indicates whether to get the maximum (true) or minimum (false) bound.
// @return The bound value.
func GetConVarBounds(conVarHandle uint64, max bool) string {
	var __retVal string
	var __retVal_native plugify.PlgString
	__conVarHandle := C.uint64_t(conVarHandle)
	__max := C.bool(max)
	plugify.Block{
		Try: func() {
			__native := C.GetConVarBounds(__conVarHandle, __max)
			__retVal_native = *(*plugify.PlgString)(unsafe.Pointer(&__native))
			// Unmarshal - Convert native data to managed data.
			__retVal = plugify.GetStringData(&__retVal_native)
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__retVal_native)
		},
	}.Do()
	return __retVal
}

// SetConVarBounds - Sets the specified bound (max or min) for a console variable.
// @param conVarHandle: The handle to the console variable data.
// @param max: Indicates whether to set the maximum (true) or minimum (false) bound.
// @param value: The value to set as the bound.
func SetConVarBounds(conVarHandle uint64, max bool, value string) {
	__conVarHandle := C.uint64_t(conVarHandle)
	__max := C.bool(max)
	__value := plugify.ConstructString(value)
	plugify.Block{
		Try: func() {
			C.SetConVarBounds(__conVarHandle, __max, (*C.String)(unsafe.Pointer(&__value)))
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__value)
		},
	}.Do()
}

// GetConVarDefault - Retrieves the default value of a console variable and stores it in the output string.
// @param conVarHandle: The handle to the console variable data.
// @return The output value in string format.
func GetConVarDefault(conVarHandle uint64) string {
	var __retVal string
	var __retVal_native plugify.PlgString
	__conVarHandle := C.uint64_t(conVarHandle)
	plugify.Block{
		Try: func() {
			__native := C.GetConVarDefault(__conVarHandle)
			__retVal_native = *(*plugify.PlgString)(unsafe.Pointer(&__native))
			// Unmarshal - Convert native data to managed data.
			__retVal = plugify.GetStringData(&__retVal_native)
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__retVal_native)
		},
	}.Do()
	return __retVal
}

// GetConVarValue - Retrieves the current value of a console variable and stores it in the output string.
// @param conVarHandle: The handle to the console variable data.
// @return The output value in string format.
func GetConVarValue(conVarHandle uint64) string {
	var __retVal string
	var __retVal_native plugify.PlgString
	__conVarHandle := C.uint64_t(conVarHandle)
	plugify.Block{
		Try: func() {
			__native := C.GetConVarValue(__conVarHandle)
			__retVal_native = *(*plugify.PlgString)(unsafe.Pointer(&__native))
			// Unmarshal - Convert native data to managed data.
			__retVal = plugify.GetStringData(&__retVal_native)
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__retVal_native)
		},
	}.Do()
	return __retVal
}

// GetConVar - Retrieves the current value of a console variable and stores it in the output.
// @param conVarHandle: The handle to the console variable data.
// @return The output value.
func GetConVar(conVarHandle uint64) interface{} {
	var __retVal interface{}
	var __retVal_native plugify.PlgVariant
	__conVarHandle := C.uint64_t(conVarHandle)
	plugify.Block{
		Try: func() {
			__native := C.GetConVar(__conVarHandle)
			__retVal_native = *(*plugify.PlgVariant)(unsafe.Pointer(&__native))
			// Unmarshal - Convert native data to managed data.
			__retVal = plugify.GetVariantData(&__retVal_native)
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyVariant(&__retVal_native)
		},
	}.Do()
	return __retVal
}

// GetConVarBool - Retrieves the current value of a boolean console variable.
// @param conVarHandle: The handle to the console variable data.
// @return The current boolean value of the console variable.
func GetConVarBool(conVarHandle uint64) bool {
	var __retVal bool
	__conVarHandle := C.uint64_t(conVarHandle)
	__retVal = bool(C.GetConVarBool(__conVarHandle))
	return __retVal
}

// GetConVarInt16 - Retrieves the current value of a signed 16-bit integer console variable.
// @param conVarHandle: The handle to the console variable data.
// @return The current int16_t value of the console variable.
func GetConVarInt16(conVarHandle uint64) int16 {
	var __retVal int16
	__conVarHandle := C.uint64_t(conVarHandle)
	__retVal = int16(C.GetConVarInt16(__conVarHandle))
	return __retVal
}

// GetConVarUInt16 - Retrieves the current value of an unsigned 16-bit integer console variable.
// @param conVarHandle: The handle to the console variable data.
// @return The current uint16_t value of the console variable.
func GetConVarUInt16(conVarHandle uint64) uint16 {
	var __retVal uint16
	__conVarHandle := C.uint64_t(conVarHandle)
	__retVal = uint16(C.GetConVarUInt16(__conVarHandle))
	return __retVal
}

// GetConVarInt32 - Retrieves the current value of a signed 32-bit integer console variable.
// @param conVarHandle: The handle to the console variable data.
// @return The current int32_t value of the console variable.
func GetConVarInt32(conVarHandle uint64) int32 {
	var __retVal int32
	__conVarHandle := C.uint64_t(conVarHandle)
	__retVal = int32(C.GetConVarInt32(__conVarHandle))
	return __retVal
}

// GetConVarUInt32 - Retrieves the current value of an unsigned 32-bit integer console variable.
// @param conVarHandle: The handle to the console variable data.
// @return The current uint32_t value of the console variable.
func GetConVarUInt32(conVarHandle uint64) uint32 {
	var __retVal uint32
	__conVarHandle := C.uint64_t(conVarHandle)
	__retVal = uint32(C.GetConVarUInt32(__conVarHandle))
	return __retVal
}

// GetConVarInt64 - Retrieves the current value of a signed 64-bit integer console variable.
// @param conVarHandle: The handle to the console variable data.
// @return The current int64_t value of the console variable.
func GetConVarInt64(conVarHandle uint64) int64 {
	var __retVal int64
	__conVarHandle := C.uint64_t(conVarHandle)
	__retVal = int64(C.GetConVarInt64(__conVarHandle))
	return __retVal
}

// GetConVarUInt64 - Retrieves the current value of an unsigned 64-bit integer console variable.
// @param conVarHandle: The handle to the console variable data.
// @return The current uint64_t value of the console variable.
func GetConVarUInt64(conVarHandle uint64) uint64 {
	var __retVal uint64
	__conVarHandle := C.uint64_t(conVarHandle)
	__retVal = uint64(C.GetConVarUInt64(__conVarHandle))
	return __retVal
}

// GetConVarFloat - Retrieves the current value of a float console variable.
// @param conVarHandle: The handle to the console variable data.
// @return The current float value of the console variable.
func GetConVarFloat(conVarHandle uint64) float32 {
	var __retVal float32
	__conVarHandle := C.uint64_t(conVarHandle)
	__retVal = float32(C.GetConVarFloat(__conVarHandle))
	return __retVal
}

// GetConVarDouble - Retrieves the current value of a double console variable.
// @param conVarHandle: The handle to the console variable data.
// @return The current double value of the console variable.
func GetConVarDouble(conVarHandle uint64) float64 {
	var __retVal float64
	__conVarHandle := C.uint64_t(conVarHandle)
	__retVal = float64(C.GetConVarDouble(__conVarHandle))
	return __retVal
}

// GetConVarString - Retrieves the current value of a string console variable.
// @param conVarHandle: The handle to the console variable data.
// @return The current string value of the console variable.
func GetConVarString(conVarHandle uint64) string {
	var __retVal string
	var __retVal_native plugify.PlgString
	__conVarHandle := C.uint64_t(conVarHandle)
	plugify.Block{
		Try: func() {
			__native := C.GetConVarString(__conVarHandle)
			__retVal_native = *(*plugify.PlgString)(unsafe.Pointer(&__native))
			// Unmarshal - Convert native data to managed data.
			__retVal = plugify.GetStringData(&__retVal_native)
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__retVal_native)
		},
	}.Do()
	return __retVal
}

// GetConVarColor - Retrieves the current value of a Color console variable.
// @param conVarHandle: The handle to the console variable data.
// @return The current Color value of the console variable.
func GetConVarColor(conVarHandle uint64) int32 {
	var __retVal int32
	__conVarHandle := C.uint64_t(conVarHandle)
	__retVal = int32(C.GetConVarColor(__conVarHandle))
	return __retVal
}

// GetConVarVector2 - Retrieves the current value of a Vector2D console variable.
// @param conVarHandle: The handle to the console variable data.
// @return The current Vector2D value of the console variable.
func GetConVarVector2(conVarHandle uint64) plugify.Vector2 {
	var __retVal plugify.Vector2
	__conVarHandle := C.uint64_t(conVarHandle)
	__native := C.GetConVarVector2(__conVarHandle)
	__retVal = *(*plugify.Vector2)(unsafe.Pointer(&__native))
	return __retVal
}

// GetConVarVector - Retrieves the current value of a Vector console variable.
// @param conVarHandle: The handle to the console variable data.
// @return The current Vector value of the console variable.
func GetConVarVector(conVarHandle uint64) plugify.Vector3 {
	var __retVal plugify.Vector3
	__conVarHandle := C.uint64_t(conVarHandle)
	__native := C.GetConVarVector(__conVarHandle)
	__retVal = *(*plugify.Vector3)(unsafe.Pointer(&__native))
	return __retVal
}

// GetConVarVector4 - Retrieves the current value of a Vector4D console variable.
// @param conVarHandle: The handle to the console variable data.
// @return The current Vector4D value of the console variable.
func GetConVarVector4(conVarHandle uint64) plugify.Vector4 {
	var __retVal plugify.Vector4
	__conVarHandle := C.uint64_t(conVarHandle)
	__native := C.GetConVarVector4(__conVarHandle)
	__retVal = *(*plugify.Vector4)(unsafe.Pointer(&__native))
	return __retVal
}

// GetConVarQAngle - Retrieves the current value of a QAngle console variable.
// @param conVarHandle: The handle to the console variable data.
// @return The current QAngle value of the console variable.
func GetConVarQAngle(conVarHandle uint64) plugify.Vector3 {
	var __retVal plugify.Vector3
	__conVarHandle := C.uint64_t(conVarHandle)
	__native := C.GetConVarQAngle(__conVarHandle)
	__retVal = *(*plugify.Vector3)(unsafe.Pointer(&__native))
	return __retVal
}

// SetConVarValue - Sets the value of a console variable.
// @param conVarHandle: The handle to the console variable data.
// @param value: The string value to set for the console variable.
// @param replicate: If set to true, the new convar value will be set on all clients. This will only work if the convar has the FCVAR_REPLICATED flag and actually exists on clients.
// @param notify: If set to true, clients will be notified that the convar has changed. This will only work if the convar has the FCVAR_NOTIFY flag.
func SetConVarValue(conVarHandle uint64, value string, replicate bool, notify bool) {
	__conVarHandle := C.uint64_t(conVarHandle)
	__value := plugify.ConstructString(value)
	__replicate := C.bool(replicate)
	__notify := C.bool(notify)
	plugify.Block{
		Try: func() {
			C.SetConVarValue(__conVarHandle, (*C.String)(unsafe.Pointer(&__value)), __replicate, __notify)
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__value)
		},
	}.Do()
}

// SetConVar - Sets the value of a console variable.
// @param conVarHandle: The handle to the console variable data.
// @param value: The value to set for the console variable.
// @param replicate: If set to true, the new convar value will be set on all clients. This will only work if the convar has the FCVAR_REPLICATED flag and actually exists on clients.
// @param notify: If set to true, clients will be notified that the convar has changed. This will only work if the convar has the FCVAR_NOTIFY flag.
func SetConVar(conVarHandle uint64, value interface{}, replicate bool, notify bool) {
	__conVarHandle := C.uint64_t(conVarHandle)
	__value := plugify.ConstructVariant(value)
	__replicate := C.bool(replicate)
	__notify := C.bool(notify)
	plugify.Block{
		Try: func() {
			C.SetConVar(__conVarHandle, (*C.Variant)(unsafe.Pointer(&__value)), __replicate, __notify)
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyVariant(&__value)
		},
	}.Do()
}

// SetConVarBool - Sets the value of a boolean console variable.
// @param conVarHandle: The handle to the console variable data.
// @param value: The value to set for the console variable.
// @param replicate: If set to true, the new convar value will be set on all clients. This will only work if the convar has the FCVAR_REPLICATED flag and actually exists on clients.
// @param notify: If set to true, clients will be notified that the convar has changed. This will only work if the convar has the FCVAR_NOTIFY flag.
func SetConVarBool(conVarHandle uint64, value bool, replicate bool, notify bool) {
	__conVarHandle := C.uint64_t(conVarHandle)
	__value := C.bool(value)
	__replicate := C.bool(replicate)
	__notify := C.bool(notify)
	C.SetConVarBool(__conVarHandle, __value, __replicate, __notify)
}

// SetConVarInt16 - Sets the value of a signed 16-bit integer console variable.
// @param conVarHandle: The handle to the console variable data.
// @param value: The value to set for the console variable.
// @param replicate: If set to true, the new convar value will be set on all clients. This will only work if the convar has the FCVAR_REPLICATED flag and actually exists on clients.
// @param notify: If set to true, clients will be notified that the convar has changed. This will only work if the convar has the FCVAR_NOTIFY flag.
func SetConVarInt16(conVarHandle uint64, value int16, replicate bool, notify bool) {
	__conVarHandle := C.uint64_t(conVarHandle)
	__value := C.int16_t(value)
	__replicate := C.bool(replicate)
	__notify := C.bool(notify)
	C.SetConVarInt16(__conVarHandle, __value, __replicate, __notify)
}

// SetConVarUInt16 - Sets the value of an unsigned 16-bit integer console variable.
// @param conVarHandle: The handle to the console variable data.
// @param value: The value to set for the console variable.
// @param replicate: If set to true, the new convar value will be set on all clients. This will only work if the convar has the FCVAR_REPLICATED flag and actually exists on clients.
// @param notify: If set to true, clients will be notified that the convar has changed. This will only work if the convar has the FCVAR_NOTIFY flag.
func SetConVarUInt16(conVarHandle uint64, value uint16, replicate bool, notify bool) {
	__conVarHandle := C.uint64_t(conVarHandle)
	__value := C.uint16_t(value)
	__replicate := C.bool(replicate)
	__notify := C.bool(notify)
	C.SetConVarUInt16(__conVarHandle, __value, __replicate, __notify)
}

// SetConVarInt32 - Sets the value of a signed 32-bit integer console variable.
// @param conVarHandle: The handle to the console variable data.
// @param value: The value to set for the console variable.
// @param replicate: If set to true, the new convar value will be set on all clients. This will only work if the convar has the FCVAR_REPLICATED flag and actually exists on clients.
// @param notify: If set to true, clients will be notified that the convar has changed. This will only work if the convar has the FCVAR_NOTIFY flag.
func SetConVarInt32(conVarHandle uint64, value int32, replicate bool, notify bool) {
	__conVarHandle := C.uint64_t(conVarHandle)
	__value := C.int32_t(value)
	__replicate := C.bool(replicate)
	__notify := C.bool(notify)
	C.SetConVarInt32(__conVarHandle, __value, __replicate, __notify)
}

// SetConVarUInt32 - Sets the value of an unsigned 32-bit integer console variable.
// @param conVarHandle: The handle to the console variable data.
// @param value: The value to set for the console variable.
// @param replicate: If set to true, the new convar value will be set on all clients. This will only work if the convar has the FCVAR_REPLICATED flag and actually exists on clients.
// @param notify: If set to true, clients will be notified that the convar has changed. This will only work if the convar has the FCVAR_NOTIFY flag.
func SetConVarUInt32(conVarHandle uint64, value uint32, replicate bool, notify bool) {
	__conVarHandle := C.uint64_t(conVarHandle)
	__value := C.uint32_t(value)
	__replicate := C.bool(replicate)
	__notify := C.bool(notify)
	C.SetConVarUInt32(__conVarHandle, __value, __replicate, __notify)
}

// SetConVarInt64 - Sets the value of a signed 64-bit integer console variable.
// @param conVarHandle: The handle to the console variable data.
// @param value: The value to set for the console variable.
// @param replicate: If set to true, the new convar value will be set on all clients. This will only work if the convar has the FCVAR_REPLICATED flag and actually exists on clients.
// @param notify: If set to true, clients will be notified that the convar has changed. This will only work if the convar has the FCVAR_NOTIFY flag.
func SetConVarInt64(conVarHandle uint64, value int64, replicate bool, notify bool) {
	__conVarHandle := C.uint64_t(conVarHandle)
	__value := C.int64_t(value)
	__replicate := C.bool(replicate)
	__notify := C.bool(notify)
	C.SetConVarInt64(__conVarHandle, __value, __replicate, __notify)
}

// SetConVarUInt64 - Sets the value of an unsigned 64-bit integer console variable.
// @param conVarHandle: The handle to the console variable data.
// @param value: The value to set for the console variable.
// @param replicate: If set to true, the new convar value will be set on all clients. This will only work if the convar has the FCVAR_REPLICATED flag and actually exists on clients.
// @param notify: If set to true, clients will be notified that the convar has changed. This will only work if the convar has the FCVAR_NOTIFY flag.
func SetConVarUInt64(conVarHandle uint64, value uint64, replicate bool, notify bool) {
	__conVarHandle := C.uint64_t(conVarHandle)
	__value := C.uint64_t(value)
	__replicate := C.bool(replicate)
	__notify := C.bool(notify)
	C.SetConVarUInt64(__conVarHandle, __value, __replicate, __notify)
}

// SetConVarFloat - Sets the value of a floating-point console variable.
// @param conVarHandle: The handle to the console variable data.
// @param value: The value to set for the console variable.
// @param replicate: If set to true, the new convar value will be set on all clients. This will only work if the convar has the FCVAR_REPLICATED flag and actually exists on clients.
// @param notify: If set to true, clients will be notified that the convar has changed. This will only work if the convar has the FCVAR_NOTIFY flag.
func SetConVarFloat(conVarHandle uint64, value float32, replicate bool, notify bool) {
	__conVarHandle := C.uint64_t(conVarHandle)
	__value := C.float(value)
	__replicate := C.bool(replicate)
	__notify := C.bool(notify)
	C.SetConVarFloat(__conVarHandle, __value, __replicate, __notify)
}

// SetConVarDouble - Sets the value of a double-precision floating-point console variable.
// @param conVarHandle: The handle to the console variable data.
// @param value: The value to set for the console variable.
// @param replicate: If set to true, the new convar value will be set on all clients. This will only work if the convar has the FCVAR_REPLICATED flag and actually exists on clients.
// @param notify: If set to true, clients will be notified that the convar has changed. This will only work if the convar has the FCVAR_NOTIFY flag.
func SetConVarDouble(conVarHandle uint64, value float64, replicate bool, notify bool) {
	__conVarHandle := C.uint64_t(conVarHandle)
	__value := C.double(value)
	__replicate := C.bool(replicate)
	__notify := C.bool(notify)
	C.SetConVarDouble(__conVarHandle, __value, __replicate, __notify)
}

// SetConVarString - Sets the value of a string console variable.
// @param conVarHandle: The handle to the console variable data.
// @param value: The value to set for the console variable.
// @param replicate: If set to true, the new convar value will be set on all clients. This will only work if the convar has the FCVAR_REPLICATED flag and actually exists on clients.
// @param notify: If set to true, clients will be notified that the convar has changed. This will only work if the convar has the FCVAR_NOTIFY flag.
func SetConVarString(conVarHandle uint64, value string, replicate bool, notify bool) {
	__conVarHandle := C.uint64_t(conVarHandle)
	__value := plugify.ConstructString(value)
	__replicate := C.bool(replicate)
	__notify := C.bool(notify)
	plugify.Block{
		Try: func() {
			C.SetConVarString(__conVarHandle, (*C.String)(unsafe.Pointer(&__value)), __replicate, __notify)
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__value)
		},
	}.Do()
}

// SetConVarColor - Sets the value of a color console variable.
// @param conVarHandle: The handle to the console variable data.
// @param value: The value to set for the console variable.
// @param replicate: If set to true, the new convar value will be set on all clients. This will only work if the convar has the FCVAR_REPLICATED flag and actually exists on clients.
// @param notify: If set to true, clients will be notified that the convar has changed. This will only work if the convar has the FCVAR_NOTIFY flag.
func SetConVarColor(conVarHandle uint64, value int32, replicate bool, notify bool) {
	__conVarHandle := C.uint64_t(conVarHandle)
	__value := C.int32_t(value)
	__replicate := C.bool(replicate)
	__notify := C.bool(notify)
	C.SetConVarColor(__conVarHandle, __value, __replicate, __notify)
}

// SetConVarVector2 - Sets the value of a 2D vector console variable.
// @param conVarHandle: The handle to the console variable data.
// @param value: The value to set for the console variable.
// @param replicate: If set to true, the new convar value will be set on all clients. This will only work if the convar has the FCVAR_REPLICATED flag and actually exists on clients.
// @param notify: If set to true, clients will be notified that the convar has changed. This will only work if the convar has the FCVAR_NOTIFY flag.
func SetConVarVector2(conVarHandle uint64, value plugify.Vector2, replicate bool, notify bool) {
	__conVarHandle := C.uint64_t(conVarHandle)
	__value := *(*C.Vector2)(unsafe.Pointer(&value))
	__replicate := C.bool(replicate)
	__notify := C.bool(notify)
	C.SetConVarVector2(__conVarHandle, &__value, __replicate, __notify)
}

// SetConVarVector3 - Sets the value of a 3D vector console variable.
// @param conVarHandle: The handle to the console variable data.
// @param value: The value to set for the console variable.
// @param replicate: If set to true, the new convar value will be set on all clients. This will only work if the convar has the FCVAR_REPLICATED flag and actually exists on clients.
// @param notify: If set to true, clients will be notified that the convar has changed. This will only work if the convar has the FCVAR_NOTIFY flag.
func SetConVarVector3(conVarHandle uint64, value plugify.Vector3, replicate bool, notify bool) {
	__conVarHandle := C.uint64_t(conVarHandle)
	__value := *(*C.Vector3)(unsafe.Pointer(&value))
	__replicate := C.bool(replicate)
	__notify := C.bool(notify)
	C.SetConVarVector3(__conVarHandle, &__value, __replicate, __notify)
}

// SetConVarVector4 - Sets the value of a 4D vector console variable.
// @param conVarHandle: The handle to the console variable data.
// @param value: The value to set for the console variable.
// @param replicate: If set to true, the new convar value will be set on all clients. This will only work if the convar has the FCVAR_REPLICATED flag and actually exists on clients.
// @param notify: If set to true, clients will be notified that the convar has changed. This will only work if the convar has the FCVAR_NOTIFY flag.
func SetConVarVector4(conVarHandle uint64, value plugify.Vector4, replicate bool, notify bool) {
	__conVarHandle := C.uint64_t(conVarHandle)
	__value := *(*C.Vector4)(unsafe.Pointer(&value))
	__replicate := C.bool(replicate)
	__notify := C.bool(notify)
	C.SetConVarVector4(__conVarHandle, &__value, __replicate, __notify)
}

// SetConVarQAngle - Sets the value of a quaternion angle console variable.
// @param conVarHandle: The handle to the console variable data.
// @param value: The value to set for the console variable.
// @param replicate: If set to true, the new convar value will be set on all clients. This will only work if the convar has the FCVAR_REPLICATED flag and actually exists on clients.
// @param notify: If set to true, clients will be notified that the convar has changed. This will only work if the convar has the FCVAR_NOTIFY flag.
func SetConVarQAngle(conVarHandle uint64, value plugify.Vector3, replicate bool, notify bool) {
	__conVarHandle := C.uint64_t(conVarHandle)
	__value := *(*C.Vector3)(unsafe.Pointer(&value))
	__replicate := C.bool(replicate)
	__notify := C.bool(notify)
	C.SetConVarQAngle(__conVarHandle, &__value, __replicate, __notify)
}

// SendConVarValue - Replicates a console variable value to a specific client. This does not change the actual console variable value.
// @param playerSlot: The index of the client to replicate the value to.
// @param conVarHandle: The handle to the console variable data.
// @param value: The value to send to the client.
func SendConVarValue(playerSlot int32, conVarHandle uint64, value string) {
	__playerSlot := C.int32_t(playerSlot)
	__conVarHandle := C.uint64_t(conVarHandle)
	__value := plugify.ConstructString(value)
	plugify.Block{
		Try: func() {
			C.SendConVarValue(__playerSlot, __conVarHandle, (*C.String)(unsafe.Pointer(&__value)))
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__value)
		},
	}.Do()
}

// GetClientConVarValue - Retrieves the value of a client's console variable and stores it in the output string.
// @param playerSlot: The index of the client whose console variable value is being retrieved.
// @param convarName: The name of the console variable to retrieve.
// @return The output string to store the client's console variable value.
func GetClientConVarValue(playerSlot int32, convarName string) string {
	var __retVal string
	var __retVal_native plugify.PlgString
	__playerSlot := C.int32_t(playerSlot)
	__convarName := plugify.ConstructString(convarName)
	plugify.Block{
		Try: func() {
			__native := C.GetClientConVarValue(__playerSlot, (*C.String)(unsafe.Pointer(&__convarName)))
			__retVal_native = *(*plugify.PlgString)(unsafe.Pointer(&__native))
			// Unmarshal - Convert native data to managed data.
			__retVal = plugify.GetStringData(&__retVal_native)
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__retVal_native)
			plugify.DestroyString(&__convarName)
		},
	}.Do()
	return __retVal
}

// SetFakeClientConVarValue - Replicates a console variable value to a specific fake client. This does not change the actual console variable value.
// @param playerSlot: The index of the fake client to replicate the value to.
// @param convarName: The name of the console variable.
// @param convarValue: The value to set for the console variable.
func SetFakeClientConVarValue(playerSlot int32, convarName string, convarValue string) {
	__playerSlot := C.int32_t(playerSlot)
	__convarName := plugify.ConstructString(convarName)
	__convarValue := plugify.ConstructString(convarValue)
	plugify.Block{
		Try: func() {
			C.SetFakeClientConVarValue(__playerSlot, (*C.String)(unsafe.Pointer(&__convarName)), (*C.String)(unsafe.Pointer(&__convarValue)))
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__convarName)
			plugify.DestroyString(&__convarValue)
		},
	}.Do()
}

// GetGameDirectory - Returns the path of the game's directory.
// @return A reference to a string where the game directory path will be stored.
func GetGameDirectory() string {
	var __retVal string
	var __retVal_native plugify.PlgString
	plugify.Block{
		Try: func() {
			__native := C.GetGameDirectory()
			__retVal_native = *(*plugify.PlgString)(unsafe.Pointer(&__native))
			// Unmarshal - Convert native data to managed data.
			__retVal = plugify.GetStringData(&__retVal_native)
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__retVal_native)
		},
	}.Do()
	return __retVal
}

// GetCurrentMap - Returns the current map name.
// @return A reference to a string where the current map name will be stored.
func GetCurrentMap() string {
	var __retVal string
	var __retVal_native plugify.PlgString
	plugify.Block{
		Try: func() {
			__native := C.GetCurrentMap()
			__retVal_native = *(*plugify.PlgString)(unsafe.Pointer(&__native))
			// Unmarshal - Convert native data to managed data.
			__retVal = plugify.GetStringData(&__retVal_native)
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__retVal_native)
		},
	}.Do()
	return __retVal
}

// IsMapValid - Returns whether a specified map is valid or not.
// @param mapname: The name of the map to check for validity.
// @return True if the map is valid, false otherwise.
func IsMapValid(mapname string) bool {
	var __retVal bool
	__mapname := plugify.ConstructString(mapname)
	plugify.Block{
		Try: func() {
			__retVal = bool(C.IsMapValid((*C.String)(unsafe.Pointer(&__mapname))))
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__mapname)
		},
	}.Do()
	return __retVal
}

// GetGameTime - Returns the game time based on the game tick.
// @return The current game time.
func GetGameTime() float32 {
	__retVal := float32(C.GetGameTime())
	return __retVal
}

// GetGameTickCount - Returns the game's internal tick count.
// @return The current tick count of the game.
func GetGameTickCount() int32 {
	__retVal := int32(C.GetGameTickCount())
	return __retVal
}

// GetGameFrameTime - Returns the time the game took processing the last frame.
// @return The frame time of the last processed frame.
func GetGameFrameTime() float32 {
	__retVal := float32(C.GetGameFrameTime())
	return __retVal
}

// GetEngineTime - Returns a high-precision time value for profiling the engine.
// @return A high-precision time value.
func GetEngineTime() float64 {
	__retVal := float64(C.GetEngineTime())
	return __retVal
}

// GetMaxClients - Returns the maximum number of clients that can connect to the server.
// @return The maximum client count, or -1 if global variables are not initialized.
func GetMaxClients() int32 {
	__retVal := int32(C.GetMaxClients())
	return __retVal
}

// PrecacheGeneric - Precaches a given generic file.
// @param model: The name of the model to be precached.
// @return An integer identifier for the generic file.
func PrecacheGeneric(model string) int32 {
	var __retVal int32
	__model := plugify.ConstructString(model)
	plugify.Block{
		Try: func() {
			__retVal = int32(C.PrecacheGeneric((*C.String)(unsafe.Pointer(&__model))))
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__model)
		},
	}.Do()
	return __retVal
}

// IsGenericPrecache - Checks if a specified generic file is precached.
// @param model: The name of the generic file to check.
// @return No description available.
func IsGenericPrecache(model string) bool {
	var __retVal bool
	__model := plugify.ConstructString(model)
	plugify.Block{
		Try: func() {
			__retVal = bool(C.IsGenericPrecache((*C.String)(unsafe.Pointer(&__model))))
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__model)
		},
	}.Do()
	return __retVal
}

// PrecacheModel - Precaches a specified model.
// @param model: The name of the model to be precached.
// @return An integer identifier for the model.
func PrecacheModel(model string) int32 {
	var __retVal int32
	__model := plugify.ConstructString(model)
	plugify.Block{
		Try: func() {
			__retVal = int32(C.PrecacheModel((*C.String)(unsafe.Pointer(&__model))))
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__model)
		},
	}.Do()
	return __retVal
}

// IsModelPrecache - Checks if a specified model is precached.
// @param model: The name of the model to check.
// @return No description available.
func IsModelPrecache(model string) bool {
	var __retVal bool
	__model := plugify.ConstructString(model)
	plugify.Block{
		Try: func() {
			__retVal = bool(C.IsModelPrecache((*C.String)(unsafe.Pointer(&__model))))
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__model)
		},
	}.Do()
	return __retVal
}

// PrecacheSound - Precaches a specified sound.
// @param sound: The name of the sound to be precached.
// @param preload: A boolean indicating if the sound should be preloaded.
// @return True if the sound is successfully precached, false otherwise.
func PrecacheSound(sound string, preload bool) bool {
	var __retVal bool
	__sound := plugify.ConstructString(sound)
	__preload := C.bool(preload)
	plugify.Block{
		Try: func() {
			__retVal = bool(C.PrecacheSound((*C.String)(unsafe.Pointer(&__sound)), __preload))
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__sound)
		},
	}.Do()
	return __retVal
}

// IsSoundPrecached - Checks if a specified sound is precached.
// @param sound: The name of the sound to check.
// @return True if the sound is precached, false otherwise.
func IsSoundPrecached(sound string) bool {
	var __retVal bool
	__sound := plugify.ConstructString(sound)
	plugify.Block{
		Try: func() {
			__retVal = bool(C.IsSoundPrecached((*C.String)(unsafe.Pointer(&__sound))))
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__sound)
		},
	}.Do()
	return __retVal
}

// PrecacheDecal - Precaches a specified decal.
// @param decal: The name of the decal to be precached.
// @param preload: A boolean indicating if the decal should be preloaded.
// @return An integer identifier for the decal.
func PrecacheDecal(decal string, preload bool) int32 {
	var __retVal int32
	__decal := plugify.ConstructString(decal)
	__preload := C.bool(preload)
	plugify.Block{
		Try: func() {
			__retVal = int32(C.PrecacheDecal((*C.String)(unsafe.Pointer(&__decal)), __preload))
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__decal)
		},
	}.Do()
	return __retVal
}

// IsDecalPrecached - Checks if a specified decal is precached.
// @param decal: The name of the decal to check.
// @return True if the decal is precached, false otherwise.
func IsDecalPrecached(decal string) bool {
	var __retVal bool
	__decal := plugify.ConstructString(decal)
	plugify.Block{
		Try: func() {
			__retVal = bool(C.IsDecalPrecached((*C.String)(unsafe.Pointer(&__decal))))
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__decal)
		},
	}.Do()
	return __retVal
}

// GetEconItemSystem - Returns a pointer to the Economy Item System.
// @return A pointer to the Econ Item System.
func GetEconItemSystem() uintptr {
	__retVal := uintptr(C.GetEconItemSystem())
	return __retVal
}

// IsServerPaused - Checks if the server is currently paused.
// @return True if the server is paused, false otherwise.
func IsServerPaused() bool {
	__retVal := bool(C.IsServerPaused())
	return __retVal
}

// QueueTaskForNextFrame - Queues a task to be executed on the next frame.
// @param callback: A callback function to be executed on the next frame.
// @param userData: An array intended to hold user-related data, allowing for elements of any type.
func QueueTaskForNextFrame(callback TaskCallback, userData []interface{}) {
	__callback := plugify.GetFunctionPointerForDelegate(callback)
	__userData := plugify.ConstructVectorVariant(userData)
	plugify.Block{
		Try: func() {
			C.QueueTaskForNextFrame(__callback, (*C.Vector)(unsafe.Pointer(&__userData)))
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyVectorVariant(&__userData)
		},
	}.Do()
}

// QueueTaskForNextWorldUpdate - Queues a task to be executed on the next world update.
// @param callback: A callback function to be executed on the next world update.
// @param userData: An array intended to hold user-related data, allowing for elements of any type.
func QueueTaskForNextWorldUpdate(callback TaskCallback, userData []interface{}) {
	__callback := plugify.GetFunctionPointerForDelegate(callback)
	__userData := plugify.ConstructVectorVariant(userData)
	plugify.Block{
		Try: func() {
			C.QueueTaskForNextWorldUpdate(__callback, (*C.Vector)(unsafe.Pointer(&__userData)))
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyVectorVariant(&__userData)
		},
	}.Do()
}

// GetSoundDuration - Returns the duration of a specified sound.
// @param name: The name of the sound to check.
// @return The duration of the sound in seconds.
func GetSoundDuration(name string) float32 {
	var __retVal float32
	__name := plugify.ConstructString(name)
	plugify.Block{
		Try: func() {
			__retVal = float32(C.GetSoundDuration((*C.String)(unsafe.Pointer(&__name))))
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__name)
		},
	}.Do()
	return __retVal
}

// EmitSound - Emits a sound from a specified entity.
// @param entityHandle: The handle of the entity that will emit the sound.
// @param sound: The name of the sound to emit.
// @param pitch: The pitch of the sound.
// @param volume: The volume of the sound.
// @param delay: The delay before the sound is played.
func EmitSound(entityHandle int32, sound string, pitch int32, volume float32, delay float32) {
	__entityHandle := C.int32_t(entityHandle)
	__sound := plugify.ConstructString(sound)
	__pitch := C.int32_t(pitch)
	__volume := C.float(volume)
	__delay := C.float(delay)
	plugify.Block{
		Try: func() {
			C.EmitSound(__entityHandle, (*C.String)(unsafe.Pointer(&__sound)), __pitch, __volume, __delay)
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__sound)
		},
	}.Do()
}

// EmitSoundToClient - Emits a sound to a specific client.
// @param playerSlot: The index of the client to whom the sound will be emitted.
// @param channel: The channel through which the sound will be played.
// @param sound: The name of the sound to emit.
// @param volume: The volume of the sound.
// @param soundLevel: The level of the sound.
// @param flags: Additional flags for sound playback.
// @param pitch: The pitch of the sound.
// @param origin: The origin of the sound in 3D space.
// @param soundTime: The time at which the sound should be played.
func EmitSoundToClient(playerSlot int32, channel int32, sound string, volume float32, soundLevel int32, flags int32, pitch int32, origin plugify.Vector3, soundTime float32) {
	__playerSlot := C.int32_t(playerSlot)
	__channel := C.int32_t(channel)
	__sound := plugify.ConstructString(sound)
	__volume := C.float(volume)
	__soundLevel := C.int32_t(soundLevel)
	__flags := C.int32_t(flags)
	__pitch := C.int32_t(pitch)
	__origin := *(*C.Vector3)(unsafe.Pointer(&origin))
	__soundTime := C.float(soundTime)
	plugify.Block{
		Try: func() {
			C.EmitSoundToClient(__playerSlot, __channel, (*C.String)(unsafe.Pointer(&__sound)), __volume, __soundLevel, __flags, __pitch, &__origin, __soundTime)
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__sound)
		},
	}.Do()
}

// EntIndexToEntPointer - Converts an entity index into an entity pointer.
// @param entityIndex: The index of the entity to convert.
// @return A pointer to the entity instance, or nullptr if the entity does not exist.
func EntIndexToEntPointer(entityIndex int32) uintptr {
	var __retVal uintptr
	__entityIndex := C.int32_t(entityIndex)
	__retVal = uintptr(C.EntIndexToEntPointer(__entityIndex))
	return __retVal
}

// EntPointerToEntIndex - Retrieves the entity index from an entity pointer.
// @param entity: A pointer to the entity whose index is to be retrieved.
// @return The index of the entity, or -1 if the entity is nullptr.
func EntPointerToEntIndex(entity uintptr) int32 {
	var __retVal int32
	__entity := C.uintptr_t(entity)
	__retVal = int32(C.EntPointerToEntIndex(__entity))
	return __retVal
}

// EntPointerToEntHandle - Converts an entity pointer into an entity handle.
// @param entity: A pointer to the entity to convert.
// @return The entity handle as an integer, or INVALID_EHANDLE_INDEX if the entity is nullptr.
func EntPointerToEntHandle(entity uintptr) int32 {
	var __retVal int32
	__entity := C.uintptr_t(entity)
	__retVal = int32(C.EntPointerToEntHandle(__entity))
	return __retVal
}

// EntHandleToEntPointer - Retrieves the entity pointer from an entity handle.
// @param entityHandle: The entity handle to convert.
// @return A pointer to the entity instance, or nullptr if the handle is invalid.
func EntHandleToEntPointer(entityHandle int32) uintptr {
	var __retVal uintptr
	__entityHandle := C.int32_t(entityHandle)
	__retVal = uintptr(C.EntHandleToEntPointer(__entityHandle))
	return __retVal
}

// EntIndexToEntHandle - Converts an entity index into an entity handle.
// @param entityIndex: The index of the entity to convert.
// @return The entity handle as an integer, or INVALID_EHANDLE_INDEX if the entity index is invalid.
func EntIndexToEntHandle(entityIndex int32) int32 {
	var __retVal int32
	__entityIndex := C.int32_t(entityIndex)
	__retVal = int32(C.EntIndexToEntHandle(__entityIndex))
	return __retVal
}

// EntHandleToEntIndex - Retrieves the entity index from an entity handle.
// @param entityHandle: The entity handle from which to retrieve the index.
// @return The index of the entity, or -1 if the handle is invalid.
func EntHandleToEntIndex(entityHandle int32) int32 {
	var __retVal int32
	__entityHandle := C.int32_t(entityHandle)
	__retVal = int32(C.EntHandleToEntIndex(__entityHandle))
	return __retVal
}

// IsValidEntHandle - Checks if the provided entity handle is valid.
// @param entityHandle: The entity handle to check.
// @return True if the entity handle is valid, false otherwise.
func IsValidEntHandle(entityHandle int32) bool {
	var __retVal bool
	__entityHandle := C.int32_t(entityHandle)
	__retVal = bool(C.IsValidEntHandle(__entityHandle))
	return __retVal
}

// IsValidEntPointer - Checks if the provided entity pointer is valid.
// @param entity: The entity pointer to check.
// @return True if the entity pointer is valid, false otherwise.
func IsValidEntPointer(entity uintptr) bool {
	var __retVal bool
	__entity := C.uintptr_t(entity)
	__retVal = bool(C.IsValidEntPointer(__entity))
	return __retVal
}

// GetFirstActiveEntity - Retrieves the pointer to the first active entity.
// @return A pointer to the first active entity.
func GetFirstActiveEntity() uintptr {
	__retVal := uintptr(C.GetFirstActiveEntity())
	return __retVal
}

// GetConcreteEntityListPointer - Retrieves a pointer to the concrete entity list.
// @return A pointer to the entity list structure.
func GetConcreteEntityListPointer() uintptr {
	__retVal := uintptr(C.GetConcreteEntityListPointer())
	return __retVal
}

// HookEntityOutput - Adds an entity output hook on a specified entity class name.
// @param szClassname: The class name of the entity to hook the output for.
// @param szOutput: The output event name to hook.
// @param callback: The callback function to invoke when the output is fired.
// @param type: Whether the hook was in post mode (after processing) or pre mode (before processing).
// @return True if the hook was successfully added, false otherwise.
func HookEntityOutput(szClassname string, szOutput string, callback HookEntityOutputCallback, type_ HookMode) bool {
	var __retVal bool
	__szClassname := plugify.ConstructString(szClassname)
	__szOutput := plugify.ConstructString(szOutput)
	__callback := plugify.GetFunctionPointerForDelegate(callback)
	__type_ := C.uint8_t(type_)
	plugify.Block{
		Try: func() {
			__retVal = bool(C.HookEntityOutput((*C.String)(unsafe.Pointer(&__szClassname)), (*C.String)(unsafe.Pointer(&__szOutput)), __callback, __type_))
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__szClassname)
			plugify.DestroyString(&__szOutput)
		},
	}.Do()
	return __retVal
}

// UnhookEntityOutput - Removes an entity output hook.
// @param szClassname: The class name of the entity from which to unhook the output.
// @param szOutput: The output event name to unhook.
// @param callback: The callback function that was previously hooked.
// @param type: Whether the hook was in post mode (after processing) or pre mode (before processing).
// @return True if the hook was successfully removed, false otherwise.
func UnhookEntityOutput(szClassname string, szOutput string, callback HookEntityOutputCallback, type_ HookMode) bool {
	var __retVal bool
	__szClassname := plugify.ConstructString(szClassname)
	__szOutput := plugify.ConstructString(szOutput)
	__callback := plugify.GetFunctionPointerForDelegate(callback)
	__type_ := C.uint8_t(type_)
	plugify.Block{
		Try: func() {
			__retVal = bool(C.UnhookEntityOutput((*C.String)(unsafe.Pointer(&__szClassname)), (*C.String)(unsafe.Pointer(&__szOutput)), __callback, __type_))
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__szClassname)
			plugify.DestroyString(&__szOutput)
		},
	}.Do()
	return __retVal
}

// FindEntityByClassname - Searches for an entity by classname.
// @param startEntity: The entity handle from which to start the search.
// @param classname: The class name of the entity to search for.
// @return The entity handle of the found entity, or INVALID_EHANDLE_INDEX if no entity is found.
func FindEntityByClassname(startEntity int32, classname string) int32 {
	var __retVal int32
	__startEntity := C.int32_t(startEntity)
	__classname := plugify.ConstructString(classname)
	plugify.Block{
		Try: func() {
			__retVal = int32(C.FindEntityByClassname(__startEntity, (*C.String)(unsafe.Pointer(&__classname))))
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__classname)
		},
	}.Do()
	return __retVal
}

// FindEntityByName - Searches for an entity by name.
// @param startEntity: The entity handle from which to start the search.
// @param name: The name of the entity to search for.
// @return The entity handle of the found entity, or INVALID_EHANDLE_INDEX if no entity is found.
func FindEntityByName(startEntity int32, name string) int32 {
	var __retVal int32
	__startEntity := C.int32_t(startEntity)
	__name := plugify.ConstructString(name)
	plugify.Block{
		Try: func() {
			__retVal = int32(C.FindEntityByName(__startEntity, (*C.String)(unsafe.Pointer(&__name))))
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__name)
		},
	}.Do()
	return __retVal
}

// CreateEntityByName - Creates an entity by string name but does not spawn it.
// @param className: The class name of the entity to create.
// @return The entity handle of the created entity, or INVALID_EHANDLE_INDEX if the entity could not be created.
func CreateEntityByName(className string) int32 {
	var __retVal int32
	__className := plugify.ConstructString(className)
	plugify.Block{
		Try: func() {
			__retVal = int32(C.CreateEntityByName((*C.String)(unsafe.Pointer(&__className))))
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__className)
		},
	}.Do()
	return __retVal
}

// DispatchSpawn - Spawns an entity into the game.
// @param entityHandle: The handle of the entity to spawn.
func DispatchSpawn(entityHandle int32) {
	__entityHandle := C.int32_t(entityHandle)
	C.DispatchSpawn(__entityHandle)
}

// DispatchSpawn2 - Spawns an entity into the game with key-value properties.
// @param entityHandle: The handle of the entity to spawn.
// @param keys: A vector of keys representing the property names to set on the entity.
// @param values: A vector of values corresponding to the keys, representing the property values to set on the entity.
func DispatchSpawn2(entityHandle int32, keys []string, values []interface{}) {
	__entityHandle := C.int32_t(entityHandle)
	__keys := plugify.ConstructVectorString(keys)
	__values := plugify.ConstructVectorVariant(values)
	plugify.Block{
		Try: func() {
			C.DispatchSpawn2(__entityHandle, (*C.Vector)(unsafe.Pointer(&__keys)), (*C.Vector)(unsafe.Pointer(&__values)))
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyVectorString(&__keys)
			plugify.DestroyVectorVariant(&__values)
		},
	}.Do()
}

// RemoveEntity - Marks an entity for deletion.
// @param entityHandle: The handle of the entity to be deleted.
func RemoveEntity(entityHandle int32) {
	__entityHandle := C.int32_t(entityHandle)
	C.RemoveEntity(__entityHandle)
}

// GetEntityClassname - Retrieves the class name of an entity.
// @param entityHandle: The handle of the entity whose class name is to be retrieved.
// @return A string where the class name will be stored.
func GetEntityClassname(entityHandle int32) string {
	var __retVal string
	var __retVal_native plugify.PlgString
	__entityHandle := C.int32_t(entityHandle)
	plugify.Block{
		Try: func() {
			__native := C.GetEntityClassname(__entityHandle)
			__retVal_native = *(*plugify.PlgString)(unsafe.Pointer(&__native))
			// Unmarshal - Convert native data to managed data.
			__retVal = plugify.GetStringData(&__retVal_native)
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__retVal_native)
		},
	}.Do()
	return __retVal
}

// GetEntityName - Retrieves the name of an entity.
// @param entityHandle: The handle of the entity whose name is to be retrieved.
// @return No description available.
func GetEntityName(entityHandle int32) string {
	var __retVal string
	var __retVal_native plugify.PlgString
	__entityHandle := C.int32_t(entityHandle)
	plugify.Block{
		Try: func() {
			__native := C.GetEntityName(__entityHandle)
			__retVal_native = *(*plugify.PlgString)(unsafe.Pointer(&__native))
			// Unmarshal - Convert native data to managed data.
			__retVal = plugify.GetStringData(&__retVal_native)
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__retVal_native)
		},
	}.Do()
	return __retVal
}

// SetEntityName - Sets the name of an entity.
// @param entityHandle: The handle of the entity whose name is to be set.
// @param name: The new name to set for the entity.
func SetEntityName(entityHandle int32, name string) {
	__entityHandle := C.int32_t(entityHandle)
	__name := plugify.ConstructString(name)
	plugify.Block{
		Try: func() {
			C.SetEntityName(__entityHandle, (*C.String)(unsafe.Pointer(&__name)))
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__name)
		},
	}.Do()
}

// GetEntityMoveType - Retrieves the movement type of an entity.
// @param entityHandle: The handle of the entity whose movement type is to be retrieved.
// @return The movement type of the entity, or 0 if the entity is invalid.
func GetEntityMoveType(entityHandle int32) MoveType {
	var __retVal MoveType
	__entityHandle := C.int32_t(entityHandle)
	__retVal = int32(C.GetEntityMoveType(__entityHandle))
	return __retVal
}

// SetEntityMoveType - Sets the movement type of an entity.
// @param entityHandle: The handle of the entity whose movement type is to be set.
// @param moveType: The new movement type to set for the entity.
func SetEntityMoveType(entityHandle int32, moveType MoveType) {
	__entityHandle := C.int32_t(entityHandle)
	__moveType := C.int32_t(moveType)
	C.SetEntityMoveType(__entityHandle, __moveType)
}

// GetEntityGravity - Retrieves the gravity scale of an entity.
// @param entityHandle: The handle of the entity whose gravity scale is to be retrieved.
// @return The gravity scale of the entity, or 0.0f if the entity is invalid.
func GetEntityGravity(entityHandle int32) float32 {
	var __retVal float32
	__entityHandle := C.int32_t(entityHandle)
	__retVal = float32(C.GetEntityGravity(__entityHandle))
	return __retVal
}

// SetEntityGravity - Sets the gravity scale of an entity.
// @param entityHandle: The handle of the entity whose gravity scale is to be set.
// @param gravity: The new gravity scale to set for the entity.
func SetEntityGravity(entityHandle int32, gravity float32) {
	__entityHandle := C.int32_t(entityHandle)
	__gravity := C.float(gravity)
	C.SetEntityGravity(__entityHandle, __gravity)
}

// GetEntityFlags - Retrieves the flags of an entity.
// @param entityHandle: The handle of the entity whose flags are to be retrieved.
// @return The flags of the entity, or 0 if the entity is invalid.
func GetEntityFlags(entityHandle int32) int32 {
	var __retVal int32
	__entityHandle := C.int32_t(entityHandle)
	__retVal = int32(C.GetEntityFlags(__entityHandle))
	return __retVal
}

// SetEntityFlags - Sets the flags of an entity.
// @param entityHandle: The handle of the entity whose flags are to be set.
// @param flags: The new flags to set for the entity.
func SetEntityFlags(entityHandle int32, flags int32) {
	__entityHandle := C.int32_t(entityHandle)
	__flags := C.int32_t(flags)
	C.SetEntityFlags(__entityHandle, __flags)
}

// GetEntityRenderColor - Retrieves the render color of an entity.
// @param entityHandle: The handle of the entity whose render color is to be retrieved.
// @return The raw color value of the entity's render color, or 0 if the entity is invalid.
func GetEntityRenderColor(entityHandle int32) int32 {
	var __retVal int32
	__entityHandle := C.int32_t(entityHandle)
	__retVal = int32(C.GetEntityRenderColor(__entityHandle))
	return __retVal
}

// SetEntityRenderColor - Sets the render color of an entity.
// @param entityHandle: The handle of the entity whose render color is to be set.
// @param color: The new raw color value to set for the entity's render color.
func SetEntityRenderColor(entityHandle int32, color int32) {
	__entityHandle := C.int32_t(entityHandle)
	__color := C.int32_t(color)
	C.SetEntityRenderColor(__entityHandle, __color)
}

// GetEntityRenderMode - Retrieves the render mode of an entity.
// @param entityHandle: The handle of the entity whose render mode is to be retrieved.
// @return The render mode of the entity, or 0 if the entity is invalid.
func GetEntityRenderMode(entityHandle int32) RenderMode {
	var __retVal RenderMode
	__entityHandle := C.int32_t(entityHandle)
	__retVal = int8(C.GetEntityRenderMode(__entityHandle))
	return __retVal
}

// SetEntityRenderMode - Sets the render mode of an entity.
// @param entityHandle: The handle of the entity whose render mode is to be set.
// @param renderMode: The new render mode to set for the entity.
func SetEntityRenderMode(entityHandle int32, renderMode RenderMode) {
	__entityHandle := C.int32_t(entityHandle)
	__renderMode := C.int8_t(renderMode)
	C.SetEntityRenderMode(__entityHandle, __renderMode)
}

// GetEntityHealth - Retrieves the health of an entity.
// @param entityHandle: The handle of the entity whose health is to be retrieved.
// @return The health of the entity, or 0 if the entity is invalid.
func GetEntityHealth(entityHandle int32) int32 {
	var __retVal int32
	__entityHandle := C.int32_t(entityHandle)
	__retVal = int32(C.GetEntityHealth(__entityHandle))
	return __retVal
}

// SetEntityHealth - Sets the health of an entity.
// @param entityHandle: The handle of the entity whose health is to be set.
// @param health: The new health value to set for the entity.
func SetEntityHealth(entityHandle int32, health int32) {
	__entityHandle := C.int32_t(entityHandle)
	__health := C.int32_t(health)
	C.SetEntityHealth(__entityHandle, __health)
}

// GetTeamEntity - Retrieves the team number of an entity.
// @param entityHandle: The handle of the entity whose team number is to be retrieved.
// @return The team number of the entity, or 0 if the entity is invalid.
func GetTeamEntity(entityHandle int32) CSTeam {
	var __retVal CSTeam
	__entityHandle := C.int32_t(entityHandle)
	__retVal = int32(C.GetTeamEntity(__entityHandle))
	return __retVal
}

// SetTeamEntity - Sets the team number of an entity.
// @param entityHandle: The handle of the entity whose team number is to be set.
// @param team: The new team number to set for the entity.
func SetTeamEntity(entityHandle int32, team CSTeam) {
	__entityHandle := C.int32_t(entityHandle)
	__team := C.int32_t(team)
	C.SetTeamEntity(__entityHandle, __team)
}

// GetEntityOwner - Retrieves the owner of an entity.
// @param entityHandle: The handle of the entity whose owner is to be retrieved.
// @return The handle of the owner entity, or INVALID_EHANDLE_INDEX if the entity is invalid.
func GetEntityOwner(entityHandle int32) int32 {
	var __retVal int32
	__entityHandle := C.int32_t(entityHandle)
	__retVal = int32(C.GetEntityOwner(__entityHandle))
	return __retVal
}

// SetEntityOwner - Sets the owner of an entity.
// @param entityHandle: The handle of the entity whose owner is to be set.
// @param ownerHandle: The handle of the new owner entity.
func SetEntityOwner(entityHandle int32, ownerHandle int32) {
	__entityHandle := C.int32_t(entityHandle)
	__ownerHandle := C.int32_t(ownerHandle)
	C.SetEntityOwner(__entityHandle, __ownerHandle)
}

// GetEntityParent - Retrieves the parent of an entity.
// @param entityHandle: The handle of the entity whose parent is to be retrieved.
// @return The handle of the parent entity, or INVALID_EHANDLE_INDEX if the entity is invalid.
func GetEntityParent(entityHandle int32) int32 {
	var __retVal int32
	__entityHandle := C.int32_t(entityHandle)
	__retVal = int32(C.GetEntityParent(__entityHandle))
	return __retVal
}

// SetEntityParent - Sets the parent of an entity.
// @param entityHandle: The handle of the entity whose parent is to be set.
// @param parentHandle: The handle of the new parent entity.
func SetEntityParent(entityHandle int32, parentHandle int32) {
	__entityHandle := C.int32_t(entityHandle)
	__parentHandle := C.int32_t(parentHandle)
	C.SetEntityParent(__entityHandle, __parentHandle)
}

// GetEntityAbsOrigin - Retrieves the absolute origin of an entity.
// @param entityHandle: The handle of the entity whose absolute origin is to be retrieved.
// @return A vector where the absolute origin will be stored.
func GetEntityAbsOrigin(entityHandle int32) plugify.Vector3 {
	var __retVal plugify.Vector3
	__entityHandle := C.int32_t(entityHandle)
	__native := C.GetEntityAbsOrigin(__entityHandle)
	__retVal = *(*plugify.Vector3)(unsafe.Pointer(&__native))
	return __retVal
}

// SetEntityAbsOrigin - Sets the absolute origin of an entity.
// @param entityHandle: The handle of the entity whose absolute origin is to be set.
// @param origin: The new absolute origin to set for the entity.
func SetEntityAbsOrigin(entityHandle int32, origin plugify.Vector3) {
	__entityHandle := C.int32_t(entityHandle)
	__origin := *(*C.Vector3)(unsafe.Pointer(&origin))
	C.SetEntityAbsOrigin(__entityHandle, &__origin)
}

// GetEntityAngRotation - Retrieves the angular rotation of an entity.
// @param entityHandle: The handle of the entity whose angular rotation is to be retrieved.
// @return A QAngle where the angular rotation will be stored.
func GetEntityAngRotation(entityHandle int32) plugify.Vector3 {
	var __retVal plugify.Vector3
	__entityHandle := C.int32_t(entityHandle)
	__native := C.GetEntityAngRotation(__entityHandle)
	__retVal = *(*plugify.Vector3)(unsafe.Pointer(&__native))
	return __retVal
}

// SetEntityAngRotation - Sets the angular rotation of an entity.
// @param entityHandle: The handle of the entity whose angular rotation is to be set.
// @param angle: The new angular rotation to set for the entity.
func SetEntityAngRotation(entityHandle int32, angle plugify.Vector3) {
	__entityHandle := C.int32_t(entityHandle)
	__angle := *(*C.Vector3)(unsafe.Pointer(&angle))
	C.SetEntityAngRotation(__entityHandle, &__angle)
}

// GetEntityAbsVelocity - Retrieves the absolute velocity of an entity.
// @param entityHandle: The handle of the entity whose absolute velocity is to be retrieved.
// @return A vector where the absolute velocity will be stored.
func GetEntityAbsVelocity(entityHandle int32) plugify.Vector3 {
	var __retVal plugify.Vector3
	__entityHandle := C.int32_t(entityHandle)
	__native := C.GetEntityAbsVelocity(__entityHandle)
	__retVal = *(*plugify.Vector3)(unsafe.Pointer(&__native))
	return __retVal
}

// SetEntityAbsVelocity - Sets the absolute velocity of an entity.
// @param entityHandle: The handle of the entity whose absolute velocity is to be set.
// @param velocity: The new absolute velocity to set for the entity.
func SetEntityAbsVelocity(entityHandle int32, velocity plugify.Vector3) {
	__entityHandle := C.int32_t(entityHandle)
	__velocity := *(*C.Vector3)(unsafe.Pointer(&velocity))
	C.SetEntityAbsVelocity(__entityHandle, &__velocity)
}

// GetEntityModel - Retrieves the model name of an entity.
// @param entityHandle: The handle of the entity whose model name is to be retrieved.
// @return A string where the model name will be stored.
func GetEntityModel(entityHandle int32) string {
	var __retVal string
	var __retVal_native plugify.PlgString
	__entityHandle := C.int32_t(entityHandle)
	plugify.Block{
		Try: func() {
			__native := C.GetEntityModel(__entityHandle)
			__retVal_native = *(*plugify.PlgString)(unsafe.Pointer(&__native))
			// Unmarshal - Convert native data to managed data.
			__retVal = plugify.GetStringData(&__retVal_native)
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__retVal_native)
		},
	}.Do()
	return __retVal
}

// SetEntityModel - Sets the model name of an entity.
// @param entityHandle: The handle of the entity whose model name is to be set.
// @param model: The new model name to set for the entity.
func SetEntityModel(entityHandle int32, model string) {
	__entityHandle := C.int32_t(entityHandle)
	__model := plugify.ConstructString(model)
	plugify.Block{
		Try: func() {
			C.SetEntityModel(__entityHandle, (*C.String)(unsafe.Pointer(&__model)))
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__model)
		},
	}.Do()
}

// GetEntityWaterLevel - Retrieves the water level of an entity.
// @param entityHandle: The handle of the entity whose water level is to be retrieved.
// @return The water level of the entity, or 0.0f if the entity is invalid.
func GetEntityWaterLevel(entityHandle int32) float32 {
	var __retVal float32
	__entityHandle := C.int32_t(entityHandle)
	__retVal = float32(C.GetEntityWaterLevel(__entityHandle))
	return __retVal
}

// GetEntityGroundEntity - Retrieves the ground entity of an entity.
// @param entityHandle: The handle of the entity whose ground entity is to be retrieved.
// @return The handle of the ground entity, or INVALID_EHANDLE_INDEX if the entity is invalid.
func GetEntityGroundEntity(entityHandle int32) int32 {
	var __retVal int32
	__entityHandle := C.int32_t(entityHandle)
	__retVal = int32(C.GetEntityGroundEntity(__entityHandle))
	return __retVal
}

// GetEntityEffects - Retrieves the effects of an entity.
// @param entityHandle: The handle of the entity whose effects are to be retrieved.
// @return The effect flags of the entity, or 0 if the entity is invalid.
func GetEntityEffects(entityHandle int32) int32 {
	var __retVal int32
	__entityHandle := C.int32_t(entityHandle)
	__retVal = int32(C.GetEntityEffects(__entityHandle))
	return __retVal
}

// TeleportEntity - Teleports an entity to a specified location and orientation.
// @param entityHandle: The handle of the entity to teleport.
// @param origin: A pointer to a Vector representing the new absolute position. Can be nullptr.
// @param angles: A pointer to a QAngle representing the new orientation. Can be nullptr.
// @param velocity: A pointer to a Vector representing the new velocity. Can be nullptr.
func TeleportEntity(entityHandle int32, origin uintptr, angles uintptr, velocity uintptr) {
	__entityHandle := C.int32_t(entityHandle)
	__origin := C.uintptr_t(origin)
	__angles := C.uintptr_t(angles)
	__velocity := C.uintptr_t(velocity)
	C.TeleportEntity(__entityHandle, __origin, __angles, __velocity)
}

// AcceptInput - Invokes a named input method on a specified entity.
// @param entityHandle: The handle of the target entity that will receive the input.
// @param inputName: The name of the input action to invoke.
// @param activatorHandle: The handle of the entity that initiated the sequence of actions.
// @param callerHandle: The handle of the entity sending this event.
// @param value: The value associated with the input action.
// @param type: The type or classification of the value.
// @param outputId: An identifier for tracking the output of this operation.
func AcceptInput(entityHandle int32, inputName string, activatorHandle int32, callerHandle int32, value interface{}, type_ FieldType, outputId int32) {
	__entityHandle := C.int32_t(entityHandle)
	__inputName := plugify.ConstructString(inputName)
	__activatorHandle := C.int32_t(activatorHandle)
	__callerHandle := C.int32_t(callerHandle)
	__value := plugify.ConstructVariant(value)
	__type_ := C.int32_t(type_)
	__outputId := C.int32_t(outputId)
	plugify.Block{
		Try: func() {
			C.AcceptInput(__entityHandle, (*C.String)(unsafe.Pointer(&__inputName)), __activatorHandle, __callerHandle, (*C.Variant)(unsafe.Pointer(&__value)), __type_, __outputId)
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__inputName)
			plugify.DestroyVariant(&__value)
		},
	}.Do()
}

// HookEvent - Creates a hook for when a game event is fired.
// @param name: The name of the event to hook.
// @param pCallback: The callback function to call when the event is fired.
// @param type: Whether the hook was in post mode (after processing) or pre mode (before processing).
// @return An integer indicating the result of the hook operation.
func HookEvent(name string, pCallback EventCallback, type_ HookMode) int32 {
	var __retVal int32
	__name := plugify.ConstructString(name)
	__pCallback := plugify.GetFunctionPointerForDelegate(pCallback)
	__type_ := C.uint8_t(type_)
	plugify.Block{
		Try: func() {
			__retVal = int32(C.HookEvent((*C.String)(unsafe.Pointer(&__name)), __pCallback, __type_))
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__name)
		},
	}.Do()
	return __retVal
}

// UnhookEvent - Removes a hook for when a game event is fired.
// @param name: The name of the event to unhook.
// @param pCallback: The callback function to remove.
// @param type: Whether the hook was in post mode (after processing) or pre mode (before processing).
// @return An integer indicating the result of the unhook operation.
func UnhookEvent(name string, pCallback EventCallback, type_ HookMode) int32 {
	var __retVal int32
	__name := plugify.ConstructString(name)
	__pCallback := plugify.GetFunctionPointerForDelegate(pCallback)
	__type_ := C.uint8_t(type_)
	plugify.Block{
		Try: func() {
			__retVal = int32(C.UnhookEvent((*C.String)(unsafe.Pointer(&__name)), __pCallback, __type_))
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__name)
		},
	}.Do()
	return __retVal
}

// CreateEvent - Creates a game event to be fired later.
// @param name: The name of the event to create.
// @param force: A boolean indicating whether to force the creation of the event.
// @return A pointer to the created EventInfo structure.
func CreateEvent(name string, force bool) uintptr {
	var __retVal uintptr
	__name := plugify.ConstructString(name)
	__force := C.bool(force)
	plugify.Block{
		Try: func() {
			__retVal = uintptr(C.CreateEvent((*C.String)(unsafe.Pointer(&__name)), __force))
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__name)
		},
	}.Do()
	return __retVal
}

// FireEvent - Fires a game event.
// @param pInfo: A pointer to the EventInfo structure containing event data.
// @param bDontBroadcast: A boolean indicating whether to broadcast the event.
func FireEvent(pInfo uintptr, bDontBroadcast bool) {
	__pInfo := C.uintptr_t(pInfo)
	__bDontBroadcast := C.bool(bDontBroadcast)
	C.FireEvent(__pInfo, __bDontBroadcast)
}

// FireEventToClient - Fires a game event to a specific client.
// @param pInfo: A pointer to the EventInfo structure containing event data.
// @param playerSlot: The index of the client to fire the event to.
func FireEventToClient(pInfo uintptr, playerSlot int32) {
	__pInfo := C.uintptr_t(pInfo)
	__playerSlot := C.int32_t(playerSlot)
	C.FireEventToClient(__pInfo, __playerSlot)
}

// CancelCreatedEvent - Cancels a previously created game event that has not been fired.
// @param pInfo: A pointer to the EventInfo structure of the event to cancel.
func CancelCreatedEvent(pInfo uintptr) {
	__pInfo := C.uintptr_t(pInfo)
	C.CancelCreatedEvent(__pInfo)
}

// GetEventBool - Retrieves the boolean value of a game event's key.
// @param pInfo: A pointer to the EventInfo structure containing event data.
// @param key: The key for which to retrieve the boolean value.
// @return The boolean value associated with the key.
func GetEventBool(pInfo uintptr, key string) bool {
	var __retVal bool
	__pInfo := C.uintptr_t(pInfo)
	__key := plugify.ConstructString(key)
	plugify.Block{
		Try: func() {
			__retVal = bool(C.GetEventBool(__pInfo, (*C.String)(unsafe.Pointer(&__key))))
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__key)
		},
	}.Do()
	return __retVal
}

// GetEventFloat - Retrieves the float value of a game event's key.
// @param pInfo: A pointer to the EventInfo structure containing event data.
// @param key: The key for which to retrieve the float value.
// @return The float value associated with the key.
func GetEventFloat(pInfo uintptr, key string) float32 {
	var __retVal float32
	__pInfo := C.uintptr_t(pInfo)
	__key := plugify.ConstructString(key)
	plugify.Block{
		Try: func() {
			__retVal = float32(C.GetEventFloat(__pInfo, (*C.String)(unsafe.Pointer(&__key))))
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__key)
		},
	}.Do()
	return __retVal
}

// GetEventInt - Retrieves the integer value of a game event's key.
// @param pInfo: A pointer to the EventInfo structure containing event data.
// @param key: The key for which to retrieve the integer value.
// @return The integer value associated with the key.
func GetEventInt(pInfo uintptr, key string) int32 {
	var __retVal int32
	__pInfo := C.uintptr_t(pInfo)
	__key := plugify.ConstructString(key)
	plugify.Block{
		Try: func() {
			__retVal = int32(C.GetEventInt(__pInfo, (*C.String)(unsafe.Pointer(&__key))))
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__key)
		},
	}.Do()
	return __retVal
}

// GetEventUInt64 - Retrieves the long integer value of a game event's key.
// @param pInfo: A pointer to the EventInfo structure containing event data.
// @param key: The key for which to retrieve the long integer value.
// @return The long integer value associated with the key.
func GetEventUInt64(pInfo uintptr, key string) uint64 {
	var __retVal uint64
	__pInfo := C.uintptr_t(pInfo)
	__key := plugify.ConstructString(key)
	plugify.Block{
		Try: func() {
			__retVal = uint64(C.GetEventUInt64(__pInfo, (*C.String)(unsafe.Pointer(&__key))))
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__key)
		},
	}.Do()
	return __retVal
}

// GetEventString - Retrieves the string value of a game event's key.
// @param pInfo: A pointer to the EventInfo structure containing event data.
// @param key: The key for which to retrieve the string value.
// @return A string where the result will be stored.
func GetEventString(pInfo uintptr, key string) string {
	var __retVal string
	var __retVal_native plugify.PlgString
	__pInfo := C.uintptr_t(pInfo)
	__key := plugify.ConstructString(key)
	plugify.Block{
		Try: func() {
			__native := C.GetEventString(__pInfo, (*C.String)(unsafe.Pointer(&__key)))
			__retVal_native = *(*plugify.PlgString)(unsafe.Pointer(&__native))
			// Unmarshal - Convert native data to managed data.
			__retVal = plugify.GetStringData(&__retVal_native)
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__retVal_native)
			plugify.DestroyString(&__key)
		},
	}.Do()
	return __retVal
}

// GetEventPtr - Retrieves the pointer value of a game event's key.
// @param pInfo: A pointer to the EventInfo structure containing event data.
// @param key: The key for which to retrieve the pointer value.
// @return The pointer value associated with the key.
func GetEventPtr(pInfo uintptr, key string) uintptr {
	var __retVal uintptr
	__pInfo := C.uintptr_t(pInfo)
	__key := plugify.ConstructString(key)
	plugify.Block{
		Try: func() {
			__retVal = uintptr(C.GetEventPtr(__pInfo, (*C.String)(unsafe.Pointer(&__key))))
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__key)
		},
	}.Do()
	return __retVal
}

// GetEventPlayerController - Retrieves the player controller address of a game event's key.
// @param pInfo: A pointer to the EventInfo structure containing event data.
// @param key: The key for which to retrieve the player controller address.
// @return A pointer to the player controller associated with the key.
func GetEventPlayerController(pInfo uintptr, key string) uintptr {
	var __retVal uintptr
	__pInfo := C.uintptr_t(pInfo)
	__key := plugify.ConstructString(key)
	plugify.Block{
		Try: func() {
			__retVal = uintptr(C.GetEventPlayerController(__pInfo, (*C.String)(unsafe.Pointer(&__key))))
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__key)
		},
	}.Do()
	return __retVal
}

// GetEventPlayerIndex - Retrieves the player index of a game event's key.
// @param pInfo: A pointer to the EventInfo structure containing event data.
// @param key: The key for which to retrieve the player index.
// @return The player index associated with the key.
func GetEventPlayerIndex(pInfo uintptr, key string) int32 {
	var __retVal int32
	__pInfo := C.uintptr_t(pInfo)
	__key := plugify.ConstructString(key)
	plugify.Block{
		Try: func() {
			__retVal = int32(C.GetEventPlayerIndex(__pInfo, (*C.String)(unsafe.Pointer(&__key))))
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__key)
		},
	}.Do()
	return __retVal
}

// GetEventPlayerPawn - Retrieves the player pawn address of a game event's key.
// @param pInfo: A pointer to the EventInfo structure containing event data.
// @param key: The key for which to retrieve the player pawn address.
// @return A pointer to the player pawn associated with the key.
func GetEventPlayerPawn(pInfo uintptr, key string) uintptr {
	var __retVal uintptr
	__pInfo := C.uintptr_t(pInfo)
	__key := plugify.ConstructString(key)
	plugify.Block{
		Try: func() {
			__retVal = uintptr(C.GetEventPlayerPawn(__pInfo, (*C.String)(unsafe.Pointer(&__key))))
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__key)
		},
	}.Do()
	return __retVal
}

// GetEventEntity - Retrieves the entity address of a game event's key.
// @param pInfo: A pointer to the EventInfo structure containing event data.
// @param key: The key for which to retrieve the entity address.
// @return A pointer to the entity associated with the key.
func GetEventEntity(pInfo uintptr, key string) uintptr {
	var __retVal uintptr
	__pInfo := C.uintptr_t(pInfo)
	__key := plugify.ConstructString(key)
	plugify.Block{
		Try: func() {
			__retVal = uintptr(C.GetEventEntity(__pInfo, (*C.String)(unsafe.Pointer(&__key))))
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__key)
		},
	}.Do()
	return __retVal
}

// GetEventEntityIndex - Retrieves the entity index of a game event's key.
// @param pInfo: A pointer to the EventInfo structure containing event data.
// @param key: The key for which to retrieve the entity index.
// @return The entity index associated with the key.
func GetEventEntityIndex(pInfo uintptr, key string) int32 {
	var __retVal int32
	__pInfo := C.uintptr_t(pInfo)
	__key := plugify.ConstructString(key)
	plugify.Block{
		Try: func() {
			__retVal = int32(C.GetEventEntityIndex(__pInfo, (*C.String)(unsafe.Pointer(&__key))))
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__key)
		},
	}.Do()
	return __retVal
}

// GetEventEntityHandle - Retrieves the entity handle of a game event's key.
// @param pInfo: A pointer to the EventInfo structure containing event data.
// @param key: The key for which to retrieve the entity handle.
// @return The entity handle associated with the key.
func GetEventEntityHandle(pInfo uintptr, key string) int32 {
	var __retVal int32
	__pInfo := C.uintptr_t(pInfo)
	__key := plugify.ConstructString(key)
	plugify.Block{
		Try: func() {
			__retVal = int32(C.GetEventEntityHandle(__pInfo, (*C.String)(unsafe.Pointer(&__key))))
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__key)
		},
	}.Do()
	return __retVal
}

// GetEventName - Retrieves the name of a game event.
// @param pInfo: A pointer to the EventInfo structure containing event data.
// @return A string where the result will be stored.
func GetEventName(pInfo uintptr) string {
	var __retVal string
	var __retVal_native plugify.PlgString
	__pInfo := C.uintptr_t(pInfo)
	plugify.Block{
		Try: func() {
			__native := C.GetEventName(__pInfo)
			__retVal_native = *(*plugify.PlgString)(unsafe.Pointer(&__native))
			// Unmarshal - Convert native data to managed data.
			__retVal = plugify.GetStringData(&__retVal_native)
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__retVal_native)
		},
	}.Do()
	return __retVal
}

// SetEventBool - Sets the boolean value of a game event's key.
// @param pInfo: A pointer to the EventInfo structure containing event data.
// @param key: The key for which to set the boolean value.
// @param value: The boolean value to set.
func SetEventBool(pInfo uintptr, key string, value bool) {
	__pInfo := C.uintptr_t(pInfo)
	__key := plugify.ConstructString(key)
	__value := C.bool(value)
	plugify.Block{
		Try: func() {
			C.SetEventBool(__pInfo, (*C.String)(unsafe.Pointer(&__key)), __value)
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__key)
		},
	}.Do()
}

// SetEventFloat - Sets the floating point value of a game event's key.
// @param pInfo: A pointer to the EventInfo structure containing event data.
// @param key: The key for which to set the float value.
// @param value: The float value to set.
func SetEventFloat(pInfo uintptr, key string, value float32) {
	__pInfo := C.uintptr_t(pInfo)
	__key := plugify.ConstructString(key)
	__value := C.float(value)
	plugify.Block{
		Try: func() {
			C.SetEventFloat(__pInfo, (*C.String)(unsafe.Pointer(&__key)), __value)
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__key)
		},
	}.Do()
}

// SetEventInt - Sets the integer value of a game event's key.
// @param pInfo: A pointer to the EventInfo structure containing event data.
// @param key: The key for which to set the integer value.
// @param value: The integer value to set.
func SetEventInt(pInfo uintptr, key string, value int32) {
	__pInfo := C.uintptr_t(pInfo)
	__key := plugify.ConstructString(key)
	__value := C.int32_t(value)
	plugify.Block{
		Try: func() {
			C.SetEventInt(__pInfo, (*C.String)(unsafe.Pointer(&__key)), __value)
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__key)
		},
	}.Do()
}

// SetEventUInt64 - Sets the long integer value of a game event's key.
// @param pInfo: A pointer to the EventInfo structure containing event data.
// @param key: The key for which to set the long integer value.
// @param value: The long integer value to set.
func SetEventUInt64(pInfo uintptr, key string, value uint64) {
	__pInfo := C.uintptr_t(pInfo)
	__key := plugify.ConstructString(key)
	__value := C.uint64_t(value)
	plugify.Block{
		Try: func() {
			C.SetEventUInt64(__pInfo, (*C.String)(unsafe.Pointer(&__key)), __value)
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__key)
		},
	}.Do()
}

// SetEventString - Sets the string value of a game event's key.
// @param pInfo: A pointer to the EventInfo structure containing event data.
// @param key: The key for which to set the string value.
// @param value: The string value to set.
func SetEventString(pInfo uintptr, key string, value string) {
	__pInfo := C.uintptr_t(pInfo)
	__key := plugify.ConstructString(key)
	__value := plugify.ConstructString(value)
	plugify.Block{
		Try: func() {
			C.SetEventString(__pInfo, (*C.String)(unsafe.Pointer(&__key)), (*C.String)(unsafe.Pointer(&__value)))
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__key)
			plugify.DestroyString(&__value)
		},
	}.Do()
}

// SetEventPtr - Sets the pointer value of a game event's key.
// @param pInfo: A pointer to the EventInfo structure containing event data.
// @param key: The key for which to set the pointer value.
// @param value: The pointer value to set.
func SetEventPtr(pInfo uintptr, key string, value uintptr) {
	__pInfo := C.uintptr_t(pInfo)
	__key := plugify.ConstructString(key)
	__value := C.uintptr_t(value)
	plugify.Block{
		Try: func() {
			C.SetEventPtr(__pInfo, (*C.String)(unsafe.Pointer(&__key)), __value)
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__key)
		},
	}.Do()
}

// SetEventPlayerController - Sets the player controller address of a game event's key.
// @param pInfo: A pointer to the EventInfo structure containing event data.
// @param key: The key for which to set the player controller address.
// @param value: A pointer to the player controller to set.
func SetEventPlayerController(pInfo uintptr, key string, value uintptr) {
	__pInfo := C.uintptr_t(pInfo)
	__key := plugify.ConstructString(key)
	__value := C.uintptr_t(value)
	plugify.Block{
		Try: func() {
			C.SetEventPlayerController(__pInfo, (*C.String)(unsafe.Pointer(&__key)), __value)
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__key)
		},
	}.Do()
}

// SetEventPlayerIndex - Sets the player index value of a game event's key.
// @param pInfo: A pointer to the EventInfo structure containing event data.
// @param key: The key for which to set the player index value.
// @param value: The player index value to set.
func SetEventPlayerIndex(pInfo uintptr, key string, value int32) {
	__pInfo := C.uintptr_t(pInfo)
	__key := plugify.ConstructString(key)
	__value := C.int32_t(value)
	plugify.Block{
		Try: func() {
			C.SetEventPlayerIndex(__pInfo, (*C.String)(unsafe.Pointer(&__key)), __value)
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__key)
		},
	}.Do()
}

// SetEventEntity - Sets the entity address of a game event's key.
// @param pInfo: A pointer to the EventInfo structure containing event data.
// @param key: The key for which to set the entity address.
// @param value: A pointer to the entity to set.
func SetEventEntity(pInfo uintptr, key string, value uintptr) {
	__pInfo := C.uintptr_t(pInfo)
	__key := plugify.ConstructString(key)
	__value := C.uintptr_t(value)
	plugify.Block{
		Try: func() {
			C.SetEventEntity(__pInfo, (*C.String)(unsafe.Pointer(&__key)), __value)
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__key)
		},
	}.Do()
}

// SetEventEntityIndex - Sets the entity index of a game event's key.
// @param pInfo: A pointer to the EventInfo structure containing event data.
// @param key: The key for which to set the entity index.
// @param value: The entity index value to set.
func SetEventEntityIndex(pInfo uintptr, key string, value int32) {
	__pInfo := C.uintptr_t(pInfo)
	__key := plugify.ConstructString(key)
	__value := C.int32_t(value)
	plugify.Block{
		Try: func() {
			C.SetEventEntityIndex(__pInfo, (*C.String)(unsafe.Pointer(&__key)), __value)
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__key)
		},
	}.Do()
}

// SetEventEntityHandle - Sets the entity handle of a game event's key.
// @param pInfo: A pointer to the EventInfo structure containing event data.
// @param key: The key for which to set the entity handle.
// @param value: The entity handle value to set.
func SetEventEntityHandle(pInfo uintptr, key string, value int32) {
	__pInfo := C.uintptr_t(pInfo)
	__key := plugify.ConstructString(key)
	__value := C.int32_t(value)
	plugify.Block{
		Try: func() {
			C.SetEventEntityHandle(__pInfo, (*C.String)(unsafe.Pointer(&__key)), __value)
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__key)
		},
	}.Do()
}

// SetEventBroadcast - Sets whether an event's broadcasting will be disabled or not.
// @param pInfo: A pointer to the EventInfo structure containing event data.
// @param dontBroadcast: A boolean indicating whether to disable broadcasting.
func SetEventBroadcast(pInfo uintptr, dontBroadcast bool) {
	__pInfo := C.uintptr_t(pInfo)
	__dontBroadcast := C.bool(dontBroadcast)
	C.SetEventBroadcast(__pInfo, __dontBroadcast)
}

// LoadEventsFromFile - Load game event descriptions from a file (e.g., "resource/gameevents.res").
// @param path: The path to the file containing event descriptions.
// @param searchAll: A boolean indicating whether to search all paths for the file.
// @return An integer indicating the result of the loading operation.
func LoadEventsFromFile(path string, searchAll bool) int32 {
	var __retVal int32
	__path := plugify.ConstructString(path)
	__searchAll := C.bool(searchAll)
	plugify.Block{
		Try: func() {
			__retVal = int32(C.LoadEventsFromFile((*C.String)(unsafe.Pointer(&__path)), __searchAll))
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__path)
		},
	}.Do()
	return __retVal
}

// CloseGameConfigFile - Closes a game configuration file.
// @param pGameConfig: A pointer to the game configuration to be closed.
func CloseGameConfigFile(pGameConfig uintptr) {
	__pGameConfig := C.uintptr_t(pGameConfig)
	C.CloseGameConfigFile(__pGameConfig)
}

// LoadGameConfigFile - Loads a game configuration file.
// @param file: The path to the game configuration file to be loaded.
// @return A pointer to the loaded CGameConfig object, or nullptr if loading fails.
func LoadGameConfigFile(file string) uintptr {
	var __retVal uintptr
	__file := plugify.ConstructString(file)
	plugify.Block{
		Try: func() {
			__retVal = uintptr(C.LoadGameConfigFile((*C.String)(unsafe.Pointer(&__file))))
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__file)
		},
	}.Do()
	return __retVal
}

// GetGameConfigPath - Retrieves the path of a game configuration.
// @param pGameConfig: A pointer to the game configuration whose path is to be retrieved.
// @return A string where the path will be stored.
func GetGameConfigPath(pGameConfig uintptr) string {
	var __retVal string
	var __retVal_native plugify.PlgString
	__pGameConfig := C.uintptr_t(pGameConfig)
	plugify.Block{
		Try: func() {
			__native := C.GetGameConfigPath(__pGameConfig)
			__retVal_native = *(*plugify.PlgString)(unsafe.Pointer(&__native))
			// Unmarshal - Convert native data to managed data.
			__retVal = plugify.GetStringData(&__retVal_native)
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__retVal_native)
		},
	}.Do()
	return __retVal
}

// GetGameConfigLibrary - Retrieves a library associated with the game configuration.
// @param pGameConfig: A pointer to the game configuration from which to retrieve the library.
// @param name: The name of the library to be retrieved.
// @return A string where the library will be stored.
func GetGameConfigLibrary(pGameConfig uintptr, name string) string {
	var __retVal string
	var __retVal_native plugify.PlgString
	__pGameConfig := C.uintptr_t(pGameConfig)
	__name := plugify.ConstructString(name)
	plugify.Block{
		Try: func() {
			__native := C.GetGameConfigLibrary(__pGameConfig, (*C.String)(unsafe.Pointer(&__name)))
			__retVal_native = *(*plugify.PlgString)(unsafe.Pointer(&__native))
			// Unmarshal - Convert native data to managed data.
			__retVal = plugify.GetStringData(&__retVal_native)
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__retVal_native)
			plugify.DestroyString(&__name)
		},
	}.Do()
	return __retVal
}

// GetGameConfigSignature - Retrieves the signature associated with the game configuration.
// @param pGameConfig: A pointer to the game configuration from which to retrieve the signature.
// @param name: The name of the signature to be retrieved.
// @return A string where the signature will be stored.
func GetGameConfigSignature(pGameConfig uintptr, name string) string {
	var __retVal string
	var __retVal_native plugify.PlgString
	__pGameConfig := C.uintptr_t(pGameConfig)
	__name := plugify.ConstructString(name)
	plugify.Block{
		Try: func() {
			__native := C.GetGameConfigSignature(__pGameConfig, (*C.String)(unsafe.Pointer(&__name)))
			__retVal_native = *(*plugify.PlgString)(unsafe.Pointer(&__native))
			// Unmarshal - Convert native data to managed data.
			__retVal = plugify.GetStringData(&__retVal_native)
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__retVal_native)
			plugify.DestroyString(&__name)
		},
	}.Do()
	return __retVal
}

// GetGameConfigSymbol - Retrieves a symbol associated with the game configuration.
// @param pGameConfig: A pointer to the game configuration from which to retrieve the symbol.
// @param name: The name of the symbol to be retrieved.
// @return A string where the symbol will be stored.
func GetGameConfigSymbol(pGameConfig uintptr, name string) string {
	var __retVal string
	var __retVal_native plugify.PlgString
	__pGameConfig := C.uintptr_t(pGameConfig)
	__name := plugify.ConstructString(name)
	plugify.Block{
		Try: func() {
			__native := C.GetGameConfigSymbol(__pGameConfig, (*C.String)(unsafe.Pointer(&__name)))
			__retVal_native = *(*plugify.PlgString)(unsafe.Pointer(&__native))
			// Unmarshal - Convert native data to managed data.
			__retVal = plugify.GetStringData(&__retVal_native)
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__retVal_native)
			plugify.DestroyString(&__name)
		},
	}.Do()
	return __retVal
}

// GetGameConfigPatch - Retrieves a patch associated with the game configuration.
// @param pGameConfig: A pointer to the game configuration from which to retrieve the patch.
// @param name: The name of the patch to be retrieved.
// @return A string where the patch will be stored.
func GetGameConfigPatch(pGameConfig uintptr, name string) string {
	var __retVal string
	var __retVal_native plugify.PlgString
	__pGameConfig := C.uintptr_t(pGameConfig)
	__name := plugify.ConstructString(name)
	plugify.Block{
		Try: func() {
			__native := C.GetGameConfigPatch(__pGameConfig, (*C.String)(unsafe.Pointer(&__name)))
			__retVal_native = *(*plugify.PlgString)(unsafe.Pointer(&__native))
			// Unmarshal - Convert native data to managed data.
			__retVal = plugify.GetStringData(&__retVal_native)
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__retVal_native)
			plugify.DestroyString(&__name)
		},
	}.Do()
	return __retVal
}

// GetGameConfigOffset - Retrieves the offset associated with a name from the game configuration.
// @param pGameConfig: A pointer to the game configuration from which to retrieve the offset.
// @param name: The name whose offset is to be retrieved.
// @return The offset associated with the specified name.
func GetGameConfigOffset(pGameConfig uintptr, name string) int32 {
	var __retVal int32
	__pGameConfig := C.uintptr_t(pGameConfig)
	__name := plugify.ConstructString(name)
	plugify.Block{
		Try: func() {
			__retVal = int32(C.GetGameConfigOffset(__pGameConfig, (*C.String)(unsafe.Pointer(&__name))))
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__name)
		},
	}.Do()
	return __retVal
}

// GetGameConfigAddress - Retrieves the address associated with a name from the game configuration.
// @param pGameConfig: A pointer to the game configuration from which to retrieve the address.
// @param name: The name whose address is to be retrieved.
// @return A pointer to the address associated with the specified name.
func GetGameConfigAddress(pGameConfig uintptr, name string) uintptr {
	var __retVal uintptr
	__pGameConfig := C.uintptr_t(pGameConfig)
	__name := plugify.ConstructString(name)
	plugify.Block{
		Try: func() {
			__retVal = uintptr(C.GetGameConfigAddress(__pGameConfig, (*C.String)(unsafe.Pointer(&__name))))
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__name)
		},
	}.Do()
	return __retVal
}

// GetGameConfigMemSig - Retrieves the memory signature associated with a name from the game configuration.
// @param pGameConfig: A pointer to the game configuration from which to retrieve the memory signature.
// @param name: The name whose memory signature is to be resolved and retrieved.
// @return A pointer to the memory signature associated with the specified name.
func GetGameConfigMemSig(pGameConfig uintptr, name string) uintptr {
	var __retVal uintptr
	__pGameConfig := C.uintptr_t(pGameConfig)
	__name := plugify.ConstructString(name)
	plugify.Block{
		Try: func() {
			__retVal = uintptr(C.GetGameConfigMemSig(__pGameConfig, (*C.String)(unsafe.Pointer(&__name))))
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__name)
		},
	}.Do()
	return __retVal
}

// RegisterLoggingChannel - Registers a new logging channel with specified properties.
// @param name: The name of the logging channel.
// @param iFlags: Flags associated with the logging channel.
// @param verbosity: The verbosity level for the logging channel.
// @param color: The color for messages logged to this channel.
// @return The ID of the newly created logging channel.
func RegisterLoggingChannel(name string, iFlags int32, verbosity LoggingVerbosity, color int32) int32 {
	var __retVal int32
	__name := plugify.ConstructString(name)
	__iFlags := C.int32_t(iFlags)
	__verbosity := C.int32_t(verbosity)
	__color := C.int32_t(color)
	plugify.Block{
		Try: func() {
			__retVal = int32(C.RegisterLoggingChannel((*C.String)(unsafe.Pointer(&__name)), __iFlags, __verbosity, __color))
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__name)
		},
	}.Do()
	return __retVal
}

// AddLoggerTagToChannel - Adds a tag to a specified logging channel.
// @param channelID: The ID of the logging channel to which the tag will be added.
// @param tagName: The name of the tag to add to the channel.
func AddLoggerTagToChannel(channelID int32, tagName string) {
	__channelID := C.int32_t(channelID)
	__tagName := plugify.ConstructString(tagName)
	plugify.Block{
		Try: func() {
			C.AddLoggerTagToChannel(__channelID, (*C.String)(unsafe.Pointer(&__tagName)))
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__tagName)
		},
	}.Do()
}

// HasLoggerTag - Checks if a specified tag exists in a logging channel.
// @param channelID: The ID of the logging channel.
// @param tag: The name of the tag to check for.
// @return True if the tag exists in the channel, otherwise false.
func HasLoggerTag(channelID int32, tag string) bool {
	var __retVal bool
	__channelID := C.int32_t(channelID)
	__tag := plugify.ConstructString(tag)
	plugify.Block{
		Try: func() {
			__retVal = bool(C.HasLoggerTag(__channelID, (*C.String)(unsafe.Pointer(&__tag))))
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__tag)
		},
	}.Do()
	return __retVal
}

// IsLoggerChannelEnabledBySeverity - Checks if a logging channel is enabled based on severity.
// @param channelID: The ID of the logging channel.
// @param severity: The severity of a logging operation.
// @return True if the channel is enabled for the specified severity, otherwise false.
func IsLoggerChannelEnabledBySeverity(channelID int32, severity LoggingSeverity) bool {
	var __retVal bool
	__channelID := C.int32_t(channelID)
	__severity := C.int32_t(severity)
	__retVal = bool(C.IsLoggerChannelEnabledBySeverity(__channelID, __severity))
	return __retVal
}

// IsLoggerChannelEnabledByVerbosity - Checks if a logging channel is enabled based on verbosity.
// @param channelID: The ID of the logging channel.
// @param verbosity: The verbosity level to check.
// @return True if the channel is enabled for the specified verbosity, otherwise false.
func IsLoggerChannelEnabledByVerbosity(channelID int32, verbosity LoggingVerbosity) bool {
	var __retVal bool
	__channelID := C.int32_t(channelID)
	__verbosity := C.int32_t(verbosity)
	__retVal = bool(C.IsLoggerChannelEnabledByVerbosity(__channelID, __verbosity))
	return __retVal
}

// GetLoggerChannelVerbosity - Retrieves the verbosity level of a logging channel.
// @param channelID: The ID of the logging channel.
// @return The verbosity level of the specified logging channel.
func GetLoggerChannelVerbosity(channelID int32) int32 {
	var __retVal int32
	__channelID := C.int32_t(channelID)
	__retVal = int32(C.GetLoggerChannelVerbosity(__channelID))
	return __retVal
}

// SetLoggerChannelVerbosity - Sets the verbosity level of a logging channel.
// @param channelID: The ID of the logging channel.
// @param verbosity: The new verbosity level to set.
func SetLoggerChannelVerbosity(channelID int32, verbosity LoggingVerbosity) {
	__channelID := C.int32_t(channelID)
	__verbosity := C.int32_t(verbosity)
	C.SetLoggerChannelVerbosity(__channelID, __verbosity)
}

// SetLoggerChannelVerbosityByName - Sets the verbosity level of a logging channel by name.
// @param channelID: The ID of the logging channel.
// @param name: The name of the logging channel.
// @param verbosity: The new verbosity level to set.
func SetLoggerChannelVerbosityByName(channelID int32, name string, verbosity LoggingVerbosity) {
	__channelID := C.int32_t(channelID)
	__name := plugify.ConstructString(name)
	__verbosity := C.int32_t(verbosity)
	plugify.Block{
		Try: func() {
			C.SetLoggerChannelVerbosityByName(__channelID, (*C.String)(unsafe.Pointer(&__name)), __verbosity)
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__name)
		},
	}.Do()
}

// SetLoggerChannelVerbosityByTag - Sets the verbosity level of a logging channel by tag.
// @param channelID: The ID of the logging channel.
// @param tag: The name of the tag.
// @param verbosity: The new verbosity level to set.
func SetLoggerChannelVerbosityByTag(channelID int32, tag string, verbosity LoggingVerbosity) {
	__channelID := C.int32_t(channelID)
	__tag := plugify.ConstructString(tag)
	__verbosity := C.int32_t(verbosity)
	plugify.Block{
		Try: func() {
			C.SetLoggerChannelVerbosityByTag(__channelID, (*C.String)(unsafe.Pointer(&__tag)), __verbosity)
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__tag)
		},
	}.Do()
}

// GetLoggerChannelColor - Retrieves the color setting of a logging channel.
// @param channelID: The ID of the logging channel.
// @return The color value of the specified logging channel.
func GetLoggerChannelColor(channelID int32) int32 {
	var __retVal int32
	__channelID := C.int32_t(channelID)
	__retVal = int32(C.GetLoggerChannelColor(__channelID))
	return __retVal
}

// SetLoggerChannelColor - Sets the color setting of a logging channel.
// @param channelID: The ID of the logging channel.
// @param color: The new color value to set for the channel.
func SetLoggerChannelColor(channelID int32, color int32) {
	__channelID := C.int32_t(channelID)
	__color := C.int32_t(color)
	C.SetLoggerChannelColor(__channelID, __color)
}

// GetLoggerChannelFlags - Retrieves the flags of a logging channel.
// @param channelID: The ID of the logging channel.
// @return The flags of the specified logging channel.
func GetLoggerChannelFlags(channelID int32) int32 {
	var __retVal int32
	__channelID := C.int32_t(channelID)
	__retVal = int32(C.GetLoggerChannelFlags(__channelID))
	return __retVal
}

// SetLoggerChannelFlags - Sets the flags of a logging channel.
// @param channelID: The ID of the logging channel.
// @param eFlags: The new flags to set for the channel.
func SetLoggerChannelFlags(channelID int32, eFlags LoggingChannelFlags) {
	__channelID := C.int32_t(channelID)
	__eFlags := C.int32_t(eFlags)
	C.SetLoggerChannelFlags(__channelID, __eFlags)
}

// Log - Logs a message to a specified channel with a severity level.
// @param channelID: The ID of the logging channel.
// @param severity: The severity level for the log message.
// @param message: The message to log.
// @return An integer indicating the result of the logging operation.
func Log(channelID int32, severity LoggingSeverity, message string) int32 {
	var __retVal int32
	__channelID := C.int32_t(channelID)
	__severity := C.int32_t(severity)
	__message := plugify.ConstructString(message)
	plugify.Block{
		Try: func() {
			__retVal = int32(C.Log(__channelID, __severity, (*C.String)(unsafe.Pointer(&__message))))
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__message)
		},
	}.Do()
	return __retVal
}

// LogColored - Logs a colored message to a specified channel with a severity level.
// @param channelID: The ID of the logging channel.
// @param severity: The severity level for the log message.
// @param color: The color for the log message.
// @param message: The message to log.
// @return An integer indicating the result of the logging operation.
func LogColored(channelID int32, severity LoggingSeverity, color int32, message string) int32 {
	var __retVal int32
	__channelID := C.int32_t(channelID)
	__severity := C.int32_t(severity)
	__color := C.int32_t(color)
	__message := plugify.ConstructString(message)
	plugify.Block{
		Try: func() {
			__retVal = int32(C.LogColored(__channelID, __severity, __color, (*C.String)(unsafe.Pointer(&__message))))
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__message)
		},
	}.Do()
	return __retVal
}

// LogFull - Logs a detailed message to a specified channel, including source code info.
// @param channelID: The ID of the logging channel.
// @param severity: The severity level for the log message.
// @param file: The file name where the log call occurred.
// @param line: The line number where the log call occurred.
// @param function: The name of the function where the log call occurred.
// @param message: The message to log.
// @return An integer indicating the result of the logging operation.
func LogFull(channelID int32, severity LoggingSeverity, file string, line int32, function string, message string) int32 {
	var __retVal int32
	__channelID := C.int32_t(channelID)
	__severity := C.int32_t(severity)
	__file := plugify.ConstructString(file)
	__line := C.int32_t(line)
	__function := plugify.ConstructString(function)
	__message := plugify.ConstructString(message)
	plugify.Block{
		Try: func() {
			__retVal = int32(C.LogFull(__channelID, __severity, (*C.String)(unsafe.Pointer(&__file)), __line, (*C.String)(unsafe.Pointer(&__function)), (*C.String)(unsafe.Pointer(&__message))))
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__file)
			plugify.DestroyString(&__function)
			plugify.DestroyString(&__message)
		},
	}.Do()
	return __retVal
}

// LogFullColored - Logs a detailed colored message to a specified channel, including source code info.
// @param channelID: The ID of the logging channel.
// @param severity: The severity level for the log message.
// @param file: The file name where the log call occurred.
// @param line: The line number where the log call occurred.
// @param function: The name of the function where the log call occurred.
// @param color: The color for the log message.
// @param message: The message to log.
// @return An integer indicating the result of the logging operation.
func LogFullColored(channelID int32, severity LoggingSeverity, file string, line int32, function string, color int32, message string) int32 {
	var __retVal int32
	__channelID := C.int32_t(channelID)
	__severity := C.int32_t(severity)
	__file := plugify.ConstructString(file)
	__line := C.int32_t(line)
	__function := plugify.ConstructString(function)
	__color := C.int32_t(color)
	__message := plugify.ConstructString(message)
	plugify.Block{
		Try: func() {
			__retVal = int32(C.LogFullColored(__channelID, __severity, (*C.String)(unsafe.Pointer(&__file)), __line, (*C.String)(unsafe.Pointer(&__function)), __color, (*C.String)(unsafe.Pointer(&__message))))
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__file)
			plugify.DestroyString(&__function)
			plugify.DestroyString(&__message)
		},
	}.Do()
	return __retVal
}

// GetSchemaOffset - Get the offset of a member in a given schema class.
// @param className: The name of the class.
// @param memberName: The name of the member whose offset is to be retrieved.
// @return The offset of the member in the class.
func GetSchemaOffset(className string, memberName string) int32 {
	var __retVal int32
	__className := plugify.ConstructString(className)
	__memberName := plugify.ConstructString(memberName)
	plugify.Block{
		Try: func() {
			__retVal = int32(C.GetSchemaOffset((*C.String)(unsafe.Pointer(&__className)), (*C.String)(unsafe.Pointer(&__memberName))))
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__className)
			plugify.DestroyString(&__memberName)
		},
	}.Do()
	return __retVal
}

// GetSchemaChainOffset - Get the offset of a chain in a given schema class.
// @param className: The name of the class.
// @return The offset of the chain entity in the class.
func GetSchemaChainOffset(className string) int32 {
	var __retVal int32
	__className := plugify.ConstructString(className)
	plugify.Block{
		Try: func() {
			__retVal = int32(C.GetSchemaChainOffset((*C.String)(unsafe.Pointer(&__className))))
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__className)
		},
	}.Do()
	return __retVal
}

// IsSchemaFieldNetworked - Check if a schema field is networked.
// @param className: The name of the class.
// @param memberName: The name of the member to check.
// @return True if the member is networked, false otherwise.
func IsSchemaFieldNetworked(className string, memberName string) bool {
	var __retVal bool
	__className := plugify.ConstructString(className)
	__memberName := plugify.ConstructString(memberName)
	plugify.Block{
		Try: func() {
			__retVal = bool(C.IsSchemaFieldNetworked((*C.String)(unsafe.Pointer(&__className)), (*C.String)(unsafe.Pointer(&__memberName))))
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__className)
			plugify.DestroyString(&__memberName)
		},
	}.Do()
	return __retVal
}

// GetSchemaClassSize - Get the size of a schema class.
// @param className: The name of the class.
// @return The size of the class in bytes, or -1 if the class is not found.
func GetSchemaClassSize(className string) int32 {
	var __retVal int32
	__className := plugify.ConstructString(className)
	plugify.Block{
		Try: func() {
			__retVal = int32(C.GetSchemaClassSize((*C.String)(unsafe.Pointer(&__className))))
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__className)
		},
	}.Do()
	return __retVal
}

// GetEntData2 - Peeks into an entity's object schema and retrieves the integer value at the given offset.
// @param entity: Pointer to the instance of the class where the value is to be set.
// @param offset: The offset of the schema to use.
// @param size: Number of bytes to write (valid values are 1, 2, 4 or 8).
// @return The integer value at the given memory location.
func GetEntData2(entity uintptr, offset int32, size int32) int64 {
	var __retVal int64
	__entity := C.uintptr_t(entity)
	__offset := C.int32_t(offset)
	__size := C.int32_t(size)
	__retVal = int64(C.GetEntData2(__entity, __offset, __size))
	return __retVal
}

// SetEntData2 - Peeks into an entity's object data and sets the integer value at the given offset.
// @param entity: Pointer to the instance of the class where the value is to be set.
// @param offset: The offset of the schema to use.
// @param value: The integer value to set.
// @param size: Number of bytes to write (valid values are 1, 2, 4 or 8).
// @param changeState: If true, change will be sent over the network.
// @param chainOffset: The offset of the chain entity in the class.
func SetEntData2(entity uintptr, offset int32, value int64, size int32, changeState bool, chainOffset int32) {
	__entity := C.uintptr_t(entity)
	__offset := C.int32_t(offset)
	__value := C.int64_t(value)
	__size := C.int32_t(size)
	__changeState := C.bool(changeState)
	__chainOffset := C.int32_t(chainOffset)
	C.SetEntData2(__entity, __offset, __value, __size, __changeState, __chainOffset)
}

// GetEntDataFloat2 - Peeks into an entity's object schema and retrieves the float value at the given offset.
// @param entity: Pointer to the instance of the class where the value is to be set.
// @param offset: The offset of the schema to use.
// @param size: Number of bytes to write (valid values are 1, 2, 4 or 8).
// @return The float value at the given memory location.
func GetEntDataFloat2(entity uintptr, offset int32, size int32) float64 {
	var __retVal float64
	__entity := C.uintptr_t(entity)
	__offset := C.int32_t(offset)
	__size := C.int32_t(size)
	__retVal = float64(C.GetEntDataFloat2(__entity, __offset, __size))
	return __retVal
}

// SetEntDataFloat2 - Peeks into an entity's object data and sets the float value at the given offset.
// @param entity: Pointer to the instance of the class where the value is to be set.
// @param offset: The offset of the schema to use.
// @param value: The float value to set.
// @param size: Number of bytes to write (valid values are 1, 2, 4 or 8).
// @param changeState: If true, change will be sent over the network.
// @param chainOffset: The offset of the chain entity in the class.
func SetEntDataFloat2(entity uintptr, offset int32, value float64, size int32, changeState bool, chainOffset int32) {
	__entity := C.uintptr_t(entity)
	__offset := C.int32_t(offset)
	__value := C.double(value)
	__size := C.int32_t(size)
	__changeState := C.bool(changeState)
	__chainOffset := C.int32_t(chainOffset)
	C.SetEntDataFloat2(__entity, __offset, __value, __size, __changeState, __chainOffset)
}

// GetEntDataString2 - Peeks into an entity's object schema and retrieves the string value at the given offset.
// @param entity: Pointer to the instance of the class where the value is to be set.
// @param offset: The offset of the schema to use.
// @return The string value at the given memory location.
func GetEntDataString2(entity uintptr, offset int32) string {
	var __retVal string
	var __retVal_native plugify.PlgString
	__entity := C.uintptr_t(entity)
	__offset := C.int32_t(offset)
	plugify.Block{
		Try: func() {
			__native := C.GetEntDataString2(__entity, __offset)
			__retVal_native = *(*plugify.PlgString)(unsafe.Pointer(&__native))
			// Unmarshal - Convert native data to managed data.
			__retVal = plugify.GetStringData(&__retVal_native)
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__retVal_native)
		},
	}.Do()
	return __retVal
}

// SetEntDataString2 - Peeks into an entity's object data and sets the string at the given offset.
// @param entity: Pointer to the instance of the class where the value is to be set.
// @param offset: The offset of the schema to use.
// @param value: The string value to set.
// @param changeState: If true, change will be sent over the network.
// @param chainOffset: The offset of the chain entity in the class.
func SetEntDataString2(entity uintptr, offset int32, value string, changeState bool, chainOffset int32) {
	__entity := C.uintptr_t(entity)
	__offset := C.int32_t(offset)
	__value := plugify.ConstructString(value)
	__changeState := C.bool(changeState)
	__chainOffset := C.int32_t(chainOffset)
	plugify.Block{
		Try: func() {
			C.SetEntDataString2(__entity, __offset, (*C.String)(unsafe.Pointer(&__value)), __changeState, __chainOffset)
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__value)
		},
	}.Do()
}

// GetEntDataVector2 - Peeks into an entity's object schema and retrieves the vector value at the given offset.
// @param entity: Pointer to the instance of the class where the value is to be set.
// @param offset: The offset of the schema to use.
// @return The vector value at the given memory location.
func GetEntDataVector2(entity uintptr, offset int32) plugify.Vector3 {
	var __retVal plugify.Vector3
	__entity := C.uintptr_t(entity)
	__offset := C.int32_t(offset)
	__native := C.GetEntDataVector2(__entity, __offset)
	__retVal = *(*plugify.Vector3)(unsafe.Pointer(&__native))
	return __retVal
}

// SetEntDataVector2 - Peeks into an entity's object data and sets the vector at the given offset.
// @param entity: Pointer to the instance of the class where the value is to be set.
// @param offset: The offset of the schema to use.
// @param value: The vector value to set.
// @param changeState: If true, change will be sent over the network.
// @param chainOffset: The offset of the chain entity in the class.
func SetEntDataVector2(entity uintptr, offset int32, value plugify.Vector3, changeState bool, chainOffset int32) {
	__entity := C.uintptr_t(entity)
	__offset := C.int32_t(offset)
	__value := *(*C.Vector3)(unsafe.Pointer(&value))
	__changeState := C.bool(changeState)
	__chainOffset := C.int32_t(chainOffset)
	C.SetEntDataVector2(__entity, __offset, &__value, __changeState, __chainOffset)
}

// GetEntDataEnt2 - Peeks into an entity's object data and retrieves the entity handle at the given offset.
// @param entity: Pointer to the instance of the class where the value is to be set.
// @param offset: The offset of the schema to use.
// @return The entity handle at the given memory location.
func GetEntDataEnt2(entity uintptr, offset int32) int32 {
	var __retVal int32
	__entity := C.uintptr_t(entity)
	__offset := C.int32_t(offset)
	__retVal = int32(C.GetEntDataEnt2(__entity, __offset))
	return __retVal
}

// SetEntDataEnt2 - Peeks into an entity's object data and sets the entity handle at the given offset.
// @param entity: Pointer to the instance of the class where the value is to be set.
// @param offset: The offset of the schema to use.
// @param value: The entity handle to set.
// @param changeState: If true, change will be sent over the network.
// @param chainOffset: The offset of the chain entity in the class.
func SetEntDataEnt2(entity uintptr, offset int32, value int32, changeState bool, chainOffset int32) {
	__entity := C.uintptr_t(entity)
	__offset := C.int32_t(offset)
	__value := C.int32_t(value)
	__changeState := C.bool(changeState)
	__chainOffset := C.int32_t(chainOffset)
	C.SetEntDataEnt2(__entity, __offset, __value, __changeState, __chainOffset)
}

// ChangeEntityState2 - Updates the networked state of a schema field for a given entity pointer.
// @param entity: Pointer to the instance of the class where the value is to be set.
// @param offset: The offset of the schema to use.
// @param chainOffset: The offset of the chain entity in the class.
func ChangeEntityState2(entity uintptr, offset int32, chainOffset int32) {
	__entity := C.uintptr_t(entity)
	__offset := C.int32_t(offset)
	__chainOffset := C.int32_t(chainOffset)
	C.ChangeEntityState2(__entity, __offset, __chainOffset)
}

// GetEntData - Peeks into an entity's object schema and retrieves the integer value at the given offset.
// @param entityHandle: The handle of the entity from which the value is to be retrieved.
// @param offset: The offset of the schema to use.
// @param size: Number of bytes to write (valid values are 1, 2, 4 or 8).
// @return The integer value at the given memory location.
func GetEntData(entityHandle int32, offset int32, size int32) int64 {
	var __retVal int64
	__entityHandle := C.int32_t(entityHandle)
	__offset := C.int32_t(offset)
	__size := C.int32_t(size)
	__retVal = int64(C.GetEntData(__entityHandle, __offset, __size))
	return __retVal
}

// SetEntData - Peeks into an entity's object data and sets the integer value at the given offset.
// @param entityHandle: The handle of the entity from which the value is to be retrieved.
// @param offset: The offset of the schema to use.
// @param value: The integer value to set.
// @param size: Number of bytes to write (valid values are 1, 2, 4 or 8).
// @param changeState: If true, change will be sent over the network.
// @param chainOffset: The offset of the chain entity in the class.
func SetEntData(entityHandle int32, offset int32, value int64, size int32, changeState bool, chainOffset int32) {
	__entityHandle := C.int32_t(entityHandle)
	__offset := C.int32_t(offset)
	__value := C.int64_t(value)
	__size := C.int32_t(size)
	__changeState := C.bool(changeState)
	__chainOffset := C.int32_t(chainOffset)
	C.SetEntData(__entityHandle, __offset, __value, __size, __changeState, __chainOffset)
}

// GetEntDataFloat - Peeks into an entity's object schema and retrieves the float value at the given offset.
// @param entityHandle: The handle of the entity from which the value is to be retrieved.
// @param offset: The offset of the schema to use.
// @param size: Number of bytes to write (valid values are 1, 2, 4 or 8).
// @return The float value at the given memory location.
func GetEntDataFloat(entityHandle int32, offset int32, size int32) float64 {
	var __retVal float64
	__entityHandle := C.int32_t(entityHandle)
	__offset := C.int32_t(offset)
	__size := C.int32_t(size)
	__retVal = float64(C.GetEntDataFloat(__entityHandle, __offset, __size))
	return __retVal
}

// SetEntDataFloat - Peeks into an entity's object data and sets the float value at the given offset.
// @param entityHandle: The handle of the entity from which the value is to be retrieved.
// @param offset: The offset of the schema to use.
// @param value: The float value to set.
// @param size: Number of bytes to write (valid values are 1, 2, 4 or 8).
// @param changeState: If true, change will be sent over the network.
// @param chainOffset: The offset of the chain entity in the class.
func SetEntDataFloat(entityHandle int32, offset int32, value float64, size int32, changeState bool, chainOffset int32) {
	__entityHandle := C.int32_t(entityHandle)
	__offset := C.int32_t(offset)
	__value := C.double(value)
	__size := C.int32_t(size)
	__changeState := C.bool(changeState)
	__chainOffset := C.int32_t(chainOffset)
	C.SetEntDataFloat(__entityHandle, __offset, __value, __size, __changeState, __chainOffset)
}

// GetEntDataString - Peeks into an entity's object schema and retrieves the string value at the given offset.
// @param entityHandle: The handle of the entity from which the value is to be retrieved.
// @param offset: The offset of the schema to use.
// @return The string value at the given memory location.
func GetEntDataString(entityHandle int32, offset int32) string {
	var __retVal string
	var __retVal_native plugify.PlgString
	__entityHandle := C.int32_t(entityHandle)
	__offset := C.int32_t(offset)
	plugify.Block{
		Try: func() {
			__native := C.GetEntDataString(__entityHandle, __offset)
			__retVal_native = *(*plugify.PlgString)(unsafe.Pointer(&__native))
			// Unmarshal - Convert native data to managed data.
			__retVal = plugify.GetStringData(&__retVal_native)
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__retVal_native)
		},
	}.Do()
	return __retVal
}

// SetEntDataString - Peeks into an entity's object data and sets the string at the given offset.
// @param entityHandle: The handle of the entity from which the value is to be retrieved.
// @param offset: The offset of the schema to use.
// @param value: The string value to set.
// @param changeState: If true, change will be sent over the network.
// @param chainOffset: The offset of the chain entity in the class.
func SetEntDataString(entityHandle int32, offset int32, value string, changeState bool, chainOffset int32) {
	__entityHandle := C.int32_t(entityHandle)
	__offset := C.int32_t(offset)
	__value := plugify.ConstructString(value)
	__changeState := C.bool(changeState)
	__chainOffset := C.int32_t(chainOffset)
	plugify.Block{
		Try: func() {
			C.SetEntDataString(__entityHandle, __offset, (*C.String)(unsafe.Pointer(&__value)), __changeState, __chainOffset)
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__value)
		},
	}.Do()
}

// GetEntDataVector - Peeks into an entity's object schema and retrieves the vector value at the given offset.
// @param entityHandle: The handle of the entity from which the value is to be retrieved.
// @param offset: The offset of the schema to use.
// @return The vector value at the given memory location.
func GetEntDataVector(entityHandle int32, offset int32) plugify.Vector3 {
	var __retVal plugify.Vector3
	__entityHandle := C.int32_t(entityHandle)
	__offset := C.int32_t(offset)
	__native := C.GetEntDataVector(__entityHandle, __offset)
	__retVal = *(*plugify.Vector3)(unsafe.Pointer(&__native))
	return __retVal
}

// SetEntDataVector - Peeks into an entity's object data and sets the vector at the given offset.
// @param entityHandle: The handle of the entity from which the value is to be retrieved.
// @param offset: The offset of the schema to use.
// @param value: The vector value to set.
// @param changeState: If true, change will be sent over the network.
// @param chainOffset: The offset of the chain entity in the class.
func SetEntDataVector(entityHandle int32, offset int32, value plugify.Vector3, changeState bool, chainOffset int32) {
	__entityHandle := C.int32_t(entityHandle)
	__offset := C.int32_t(offset)
	__value := *(*C.Vector3)(unsafe.Pointer(&value))
	__changeState := C.bool(changeState)
	__chainOffset := C.int32_t(chainOffset)
	C.SetEntDataVector(__entityHandle, __offset, &__value, __changeState, __chainOffset)
}

// GetEntDataEnt - Peeks into an entity's object data and retrieves the entity handle at the given offset.
// @param entityHandle: The handle of the entity from which the value is to be retrieved.
// @param offset: The offset of the schema to use.
// @return The entity handle at the given memory location.
func GetEntDataEnt(entityHandle int32, offset int32) int32 {
	var __retVal int32
	__entityHandle := C.int32_t(entityHandle)
	__offset := C.int32_t(offset)
	__retVal = int32(C.GetEntDataEnt(__entityHandle, __offset))
	return __retVal
}

// SetEntDataEnt - Peeks into an entity's object data and sets the entity handle at the given offset.
// @param entityHandle: The handle of the entity from which the value is to be retrieved.
// @param offset: The offset of the schema to use.
// @param value: The entity handle to set.
// @param changeState: If true, change will be sent over the network.
// @param chainOffset: The offset of the chain entity in the class.
func SetEntDataEnt(entityHandle int32, offset int32, value int32, changeState bool, chainOffset int32) {
	__entityHandle := C.int32_t(entityHandle)
	__offset := C.int32_t(offset)
	__value := C.int32_t(value)
	__changeState := C.bool(changeState)
	__chainOffset := C.int32_t(chainOffset)
	C.SetEntDataEnt(__entityHandle, __offset, __value, __changeState, __chainOffset)
}

// ChangeEntityState - Updates the networked state of a schema field for a given entity handle.
// @param entityHandle: The handle of the entity from which the value is to be retrieved.
// @param offset: The offset of the schema to use.
// @param chainOffset: The offset of the chain entity in the class.
func ChangeEntityState(entityHandle int32, offset int32, chainOffset int32) {
	__entityHandle := C.int32_t(entityHandle)
	__offset := C.int32_t(offset)
	__chainOffset := C.int32_t(chainOffset)
	C.ChangeEntityState(__entityHandle, __offset, __chainOffset)
}

// GetEntSchemaArraySize2 - Retrieves the count of values that an entity schema's array can store.
// @param entity: Pointer to the instance of the class where the value is to be set.
// @param className: The name of the class.
// @param memberName: The name of the schema member.
// @return Size of array (in elements) or 0 if schema is not an array.
func GetEntSchemaArraySize2(entity uintptr, className string, memberName string) int32 {
	var __retVal int32
	__entity := C.uintptr_t(entity)
	__className := plugify.ConstructString(className)
	__memberName := plugify.ConstructString(memberName)
	plugify.Block{
		Try: func() {
			__retVal = int32(C.GetEntSchemaArraySize2(__entity, (*C.String)(unsafe.Pointer(&__className)), (*C.String)(unsafe.Pointer(&__memberName))))
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__className)
			plugify.DestroyString(&__memberName)
		},
	}.Do()
	return __retVal
}

// GetEntSchema2 - Retrieves an integer value from an entity's schema.
// @param entity: Pointer to the instance of the class where the value is to be set.
// @param className: The name of the class.
// @param memberName: The name of the schema member.
// @param element: Element # (starting from 0) if schema is an array.
// @return An integer value at the given schema offset.
func GetEntSchema2(entity uintptr, className string, memberName string, element int32) int64 {
	var __retVal int64
	__entity := C.uintptr_t(entity)
	__className := plugify.ConstructString(className)
	__memberName := plugify.ConstructString(memberName)
	__element := C.int32_t(element)
	plugify.Block{
		Try: func() {
			__retVal = int64(C.GetEntSchema2(__entity, (*C.String)(unsafe.Pointer(&__className)), (*C.String)(unsafe.Pointer(&__memberName)), __element))
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__className)
			plugify.DestroyString(&__memberName)
		},
	}.Do()
	return __retVal
}

// SetEntSchema2 - Sets an integer value in an entity's schema.
// @param entity: Pointer to the instance of the class where the value is to be set.
// @param className: The name of the class.
// @param memberName: The name of the schema member.
// @param value: The integer value to set.
// @param changeState: If true, change will be sent over the network.
// @param element: Element # (starting from 0) if schema is an array.
func SetEntSchema2(entity uintptr, className string, memberName string, value int64, changeState bool, element int32) {
	__entity := C.uintptr_t(entity)
	__className := plugify.ConstructString(className)
	__memberName := plugify.ConstructString(memberName)
	__value := C.int64_t(value)
	__changeState := C.bool(changeState)
	__element := C.int32_t(element)
	plugify.Block{
		Try: func() {
			C.SetEntSchema2(__entity, (*C.String)(unsafe.Pointer(&__className)), (*C.String)(unsafe.Pointer(&__memberName)), __value, __changeState, __element)
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__className)
			plugify.DestroyString(&__memberName)
		},
	}.Do()
}

// GetEntSchemaFloat2 - Retrieves a float value from an entity's schema.
// @param entity: Pointer to the instance of the class where the value is to be set.
// @param className: The name of the class.
// @param memberName: The name of the schema member.
// @param element: Element # (starting from 0) if schema is an array.
// @return A float value at the given schema offset.
func GetEntSchemaFloat2(entity uintptr, className string, memberName string, element int32) float64 {
	var __retVal float64
	__entity := C.uintptr_t(entity)
	__className := plugify.ConstructString(className)
	__memberName := plugify.ConstructString(memberName)
	__element := C.int32_t(element)
	plugify.Block{
		Try: func() {
			__retVal = float64(C.GetEntSchemaFloat2(__entity, (*C.String)(unsafe.Pointer(&__className)), (*C.String)(unsafe.Pointer(&__memberName)), __element))
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__className)
			plugify.DestroyString(&__memberName)
		},
	}.Do()
	return __retVal
}

// SetEntSchemaFloat2 - Sets a float value in an entity's schema.
// @param entity: Pointer to the instance of the class where the value is to be set.
// @param className: The name of the class.
// @param memberName: The name of the schema member.
// @param value: The float value to set.
// @param changeState: If true, change will be sent over the network.
// @param element: Element # (starting from 0) if schema is an array.
func SetEntSchemaFloat2(entity uintptr, className string, memberName string, value float64, changeState bool, element int32) {
	__entity := C.uintptr_t(entity)
	__className := plugify.ConstructString(className)
	__memberName := plugify.ConstructString(memberName)
	__value := C.double(value)
	__changeState := C.bool(changeState)
	__element := C.int32_t(element)
	plugify.Block{
		Try: func() {
			C.SetEntSchemaFloat2(__entity, (*C.String)(unsafe.Pointer(&__className)), (*C.String)(unsafe.Pointer(&__memberName)), __value, __changeState, __element)
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__className)
			plugify.DestroyString(&__memberName)
		},
	}.Do()
}

// GetEntSchemaString2 - Retrieves a string value from an entity's schema.
// @param entity: Pointer to the instance of the class where the value is to be set.
// @param className: The name of the class.
// @param memberName: The name of the schema member.
// @param element: Element # (starting from 0) if schema is an array.
// @return A string value at the given schema offset.
func GetEntSchemaString2(entity uintptr, className string, memberName string, element int32) string {
	var __retVal string
	var __retVal_native plugify.PlgString
	__entity := C.uintptr_t(entity)
	__className := plugify.ConstructString(className)
	__memberName := plugify.ConstructString(memberName)
	__element := C.int32_t(element)
	plugify.Block{
		Try: func() {
			__native := C.GetEntSchemaString2(__entity, (*C.String)(unsafe.Pointer(&__className)), (*C.String)(unsafe.Pointer(&__memberName)), __element)
			__retVal_native = *(*plugify.PlgString)(unsafe.Pointer(&__native))
			// Unmarshal - Convert native data to managed data.
			__retVal = plugify.GetStringData(&__retVal_native)
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__retVal_native)
			plugify.DestroyString(&__className)
			plugify.DestroyString(&__memberName)
		},
	}.Do()
	return __retVal
}

// SetEntSchemaString2 - Sets a string value in an entity's schema.
// @param entity: Pointer to the instance of the class where the value is to be set.
// @param className: The name of the class.
// @param memberName: The name of the schema member.
// @param value: The string value to set.
// @param changeState: If true, change will be sent over the network.
// @param element: Element # (starting from 0) if schema is an array.
func SetEntSchemaString2(entity uintptr, className string, memberName string, value string, changeState bool, element int32) {
	__entity := C.uintptr_t(entity)
	__className := plugify.ConstructString(className)
	__memberName := plugify.ConstructString(memberName)
	__value := plugify.ConstructString(value)
	__changeState := C.bool(changeState)
	__element := C.int32_t(element)
	plugify.Block{
		Try: func() {
			C.SetEntSchemaString2(__entity, (*C.String)(unsafe.Pointer(&__className)), (*C.String)(unsafe.Pointer(&__memberName)), (*C.String)(unsafe.Pointer(&__value)), __changeState, __element)
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__className)
			plugify.DestroyString(&__memberName)
			plugify.DestroyString(&__value)
		},
	}.Do()
}

// GetEntSchemaVector3D2 - Retrieves a vector value from an entity's schema.
// @param entity: Pointer to the instance of the class where the value is to be set.
// @param className: The name of the class.
// @param memberName: The name of the schema member.
// @param element: Element # (starting from 0) if schema is an array.
// @return A vector value at the given schema offset.
func GetEntSchemaVector3D2(entity uintptr, className string, memberName string, element int32) plugify.Vector3 {
	var __retVal plugify.Vector3
	__entity := C.uintptr_t(entity)
	__className := plugify.ConstructString(className)
	__memberName := plugify.ConstructString(memberName)
	__element := C.int32_t(element)
	plugify.Block{
		Try: func() {
			__native := C.GetEntSchemaVector3D2(__entity, (*C.String)(unsafe.Pointer(&__className)), (*C.String)(unsafe.Pointer(&__memberName)), __element)
			__retVal = *(*plugify.Vector3)(unsafe.Pointer(&__native))
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__className)
			plugify.DestroyString(&__memberName)
		},
	}.Do()
	return __retVal
}

// SetEntSchemaVector3D2 - Sets a vector value in an entity's schema.
// @param entity: Pointer to the instance of the class where the value is to be set.
// @param className: The name of the class.
// @param memberName: The name of the schema member.
// @param value: The vector value to set.
// @param changeState: If true, change will be sent over the network.
// @param element: Element # (starting from 0) if schema is an array.
func SetEntSchemaVector3D2(entity uintptr, className string, memberName string, value plugify.Vector3, changeState bool, element int32) {
	__entity := C.uintptr_t(entity)
	__className := plugify.ConstructString(className)
	__memberName := plugify.ConstructString(memberName)
	__value := *(*C.Vector3)(unsafe.Pointer(&value))
	__changeState := C.bool(changeState)
	__element := C.int32_t(element)
	plugify.Block{
		Try: func() {
			C.SetEntSchemaVector3D2(__entity, (*C.String)(unsafe.Pointer(&__className)), (*C.String)(unsafe.Pointer(&__memberName)), &__value, __changeState, __element)
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__className)
			plugify.DestroyString(&__memberName)
		},
	}.Do()
}

// GetEntSchemaVector2D2 - Retrieves a vector value from an entity's schema.
// @param entity: Pointer to the instance of the class where the value is to be set.
// @param className: The name of the class.
// @param memberName: The name of the schema member.
// @param element: Element # (starting from 0) if schema is an array.
// @return A vector value at the given schema offset.
func GetEntSchemaVector2D2(entity uintptr, className string, memberName string, element int32) plugify.Vector2 {
	var __retVal plugify.Vector2
	__entity := C.uintptr_t(entity)
	__className := plugify.ConstructString(className)
	__memberName := plugify.ConstructString(memberName)
	__element := C.int32_t(element)
	plugify.Block{
		Try: func() {
			__native := C.GetEntSchemaVector2D2(__entity, (*C.String)(unsafe.Pointer(&__className)), (*C.String)(unsafe.Pointer(&__memberName)), __element)
			__retVal = *(*plugify.Vector2)(unsafe.Pointer(&__native))
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__className)
			plugify.DestroyString(&__memberName)
		},
	}.Do()
	return __retVal
}

// SetEntSchemaVector2D2 - Sets a vector value in an entity's schema.
// @param entity: Pointer to the instance of the class where the value is to be set.
// @param className: The name of the class.
// @param memberName: The name of the schema member.
// @param value: The vector value to set.
// @param changeState: If true, change will be sent over the network.
// @param element: Element # (starting from 0) if schema is an array.
func SetEntSchemaVector2D2(entity uintptr, className string, memberName string, value plugify.Vector2, changeState bool, element int32) {
	__entity := C.uintptr_t(entity)
	__className := plugify.ConstructString(className)
	__memberName := plugify.ConstructString(memberName)
	__value := *(*C.Vector2)(unsafe.Pointer(&value))
	__changeState := C.bool(changeState)
	__element := C.int32_t(element)
	plugify.Block{
		Try: func() {
			C.SetEntSchemaVector2D2(__entity, (*C.String)(unsafe.Pointer(&__className)), (*C.String)(unsafe.Pointer(&__memberName)), &__value, __changeState, __element)
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__className)
			plugify.DestroyString(&__memberName)
		},
	}.Do()
}

// GetEntSchemaVector4D2 - Retrieves a vector value from an entity's schema.
// @param entity: Pointer to the instance of the class where the value is to be set.
// @param className: The name of the class.
// @param memberName: The name of the schema member.
// @param element: Element # (starting from 0) if schema is an array.
// @return A vector value at the given schema offset.
func GetEntSchemaVector4D2(entity uintptr, className string, memberName string, element int32) plugify.Vector4 {
	var __retVal plugify.Vector4
	__entity := C.uintptr_t(entity)
	__className := plugify.ConstructString(className)
	__memberName := plugify.ConstructString(memberName)
	__element := C.int32_t(element)
	plugify.Block{
		Try: func() {
			__native := C.GetEntSchemaVector4D2(__entity, (*C.String)(unsafe.Pointer(&__className)), (*C.String)(unsafe.Pointer(&__memberName)), __element)
			__retVal = *(*plugify.Vector4)(unsafe.Pointer(&__native))
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__className)
			plugify.DestroyString(&__memberName)
		},
	}.Do()
	return __retVal
}

// SetEntSchemaVector4D2 - Sets a vector value in an entity's schema.
// @param entity: Pointer to the instance of the class where the value is to be set.
// @param className: The name of the class.
// @param memberName: The name of the schema member.
// @param value: The vector value to set.
// @param changeState: If true, change will be sent over the network.
// @param element: Element # (starting from 0) if schema is an array.
func SetEntSchemaVector4D2(entity uintptr, className string, memberName string, value plugify.Vector4, changeState bool, element int32) {
	__entity := C.uintptr_t(entity)
	__className := plugify.ConstructString(className)
	__memberName := plugify.ConstructString(memberName)
	__value := *(*C.Vector4)(unsafe.Pointer(&value))
	__changeState := C.bool(changeState)
	__element := C.int32_t(element)
	plugify.Block{
		Try: func() {
			C.SetEntSchemaVector4D2(__entity, (*C.String)(unsafe.Pointer(&__className)), (*C.String)(unsafe.Pointer(&__memberName)), &__value, __changeState, __element)
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__className)
			plugify.DestroyString(&__memberName)
		},
	}.Do()
}

// GetEntSchemaEnt2 - Retrieves an entity handle from an entity's schema.
// @param entity: Pointer to the instance of the class where the value is to be set.
// @param className: The name of the class.
// @param memberName: The name of the schema member.
// @param element: Element # (starting from 0) if schema is an array.
// @return A string value at the given schema offset.
func GetEntSchemaEnt2(entity uintptr, className string, memberName string, element int32) int32 {
	var __retVal int32
	__entity := C.uintptr_t(entity)
	__className := plugify.ConstructString(className)
	__memberName := plugify.ConstructString(memberName)
	__element := C.int32_t(element)
	plugify.Block{
		Try: func() {
			__retVal = int32(C.GetEntSchemaEnt2(__entity, (*C.String)(unsafe.Pointer(&__className)), (*C.String)(unsafe.Pointer(&__memberName)), __element))
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__className)
			plugify.DestroyString(&__memberName)
		},
	}.Do()
	return __retVal
}

// SetEntSchemaEnt2 - Sets an entity handle in an entity's schema.
// @param entity: Pointer to the instance of the class where the value is to be set.
// @param className: The name of the class.
// @param memberName: The name of the schema member.
// @param value: The entity handle to set.
// @param changeState: If true, change will be sent over the network.
// @param element: Element # (starting from 0) if schema is an array.
func SetEntSchemaEnt2(entity uintptr, className string, memberName string, value int32, changeState bool, element int32) {
	__entity := C.uintptr_t(entity)
	__className := plugify.ConstructString(className)
	__memberName := plugify.ConstructString(memberName)
	__value := C.int32_t(value)
	__changeState := C.bool(changeState)
	__element := C.int32_t(element)
	plugify.Block{
		Try: func() {
			C.SetEntSchemaEnt2(__entity, (*C.String)(unsafe.Pointer(&__className)), (*C.String)(unsafe.Pointer(&__memberName)), __value, __changeState, __element)
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__className)
			plugify.DestroyString(&__memberName)
		},
	}.Do()
}

// NetworkStateChanged2 - Updates the networked state of a schema field for a given entity pointer.
// @param entity: Pointer to the instance of the class where the value is to be set.
// @param className: The name of the class that contains the member.
// @param memberName: The name of the member to be set.
func NetworkStateChanged2(entity uintptr, className string, memberName string) {
	__entity := C.uintptr_t(entity)
	__className := plugify.ConstructString(className)
	__memberName := plugify.ConstructString(memberName)
	plugify.Block{
		Try: func() {
			C.NetworkStateChanged2(__entity, (*C.String)(unsafe.Pointer(&__className)), (*C.String)(unsafe.Pointer(&__memberName)))
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__className)
			plugify.DestroyString(&__memberName)
		},
	}.Do()
}

// GetEntSchemaArraySize - Retrieves the count of values that an entity schema's array can store.
// @param entityHandle: The handle of the entity from which the value is to be retrieved.
// @param className: The name of the class.
// @param memberName: The name of the schema member.
// @return Size of array (in elements) or 0 if schema is not an array.
func GetEntSchemaArraySize(entityHandle int32, className string, memberName string) int32 {
	var __retVal int32
	__entityHandle := C.int32_t(entityHandle)
	__className := plugify.ConstructString(className)
	__memberName := plugify.ConstructString(memberName)
	plugify.Block{
		Try: func() {
			__retVal = int32(C.GetEntSchemaArraySize(__entityHandle, (*C.String)(unsafe.Pointer(&__className)), (*C.String)(unsafe.Pointer(&__memberName))))
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__className)
			plugify.DestroyString(&__memberName)
		},
	}.Do()
	return __retVal
}

// GetEntSchema - Retrieves an integer value from an entity's schema.
// @param entityHandle: The handle of the entity from which the value is to be retrieved.
// @param className: The name of the class.
// @param memberName: The name of the schema member.
// @param element: Element # (starting from 0) if schema is an array.
// @return An integer value at the given schema offset.
func GetEntSchema(entityHandle int32, className string, memberName string, element int32) int64 {
	var __retVal int64
	__entityHandle := C.int32_t(entityHandle)
	__className := plugify.ConstructString(className)
	__memberName := plugify.ConstructString(memberName)
	__element := C.int32_t(element)
	plugify.Block{
		Try: func() {
			__retVal = int64(C.GetEntSchema(__entityHandle, (*C.String)(unsafe.Pointer(&__className)), (*C.String)(unsafe.Pointer(&__memberName)), __element))
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__className)
			plugify.DestroyString(&__memberName)
		},
	}.Do()
	return __retVal
}

// SetEntSchema - Sets an integer value in an entity's schema.
// @param entityHandle: The handle of the entity from which the value is to be retrieved.
// @param className: The name of the class.
// @param memberName: The name of the schema member.
// @param value: The integer value to set.
// @param changeState: If true, change will be sent over the network.
// @param element: Element # (starting from 0) if schema is an array.
func SetEntSchema(entityHandle int32, className string, memberName string, value int64, changeState bool, element int32) {
	__entityHandle := C.int32_t(entityHandle)
	__className := plugify.ConstructString(className)
	__memberName := plugify.ConstructString(memberName)
	__value := C.int64_t(value)
	__changeState := C.bool(changeState)
	__element := C.int32_t(element)
	plugify.Block{
		Try: func() {
			C.SetEntSchema(__entityHandle, (*C.String)(unsafe.Pointer(&__className)), (*C.String)(unsafe.Pointer(&__memberName)), __value, __changeState, __element)
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__className)
			plugify.DestroyString(&__memberName)
		},
	}.Do()
}

// GetEntSchemaFloat - Retrieves a float value from an entity's schema.
// @param entityHandle: The handle of the entity from which the value is to be retrieved.
// @param className: The name of the class.
// @param memberName: The name of the schema member.
// @param element: Element # (starting from 0) if schema is an array.
// @return A float value at the given schema offset.
func GetEntSchemaFloat(entityHandle int32, className string, memberName string, element int32) float64 {
	var __retVal float64
	__entityHandle := C.int32_t(entityHandle)
	__className := plugify.ConstructString(className)
	__memberName := plugify.ConstructString(memberName)
	__element := C.int32_t(element)
	plugify.Block{
		Try: func() {
			__retVal = float64(C.GetEntSchemaFloat(__entityHandle, (*C.String)(unsafe.Pointer(&__className)), (*C.String)(unsafe.Pointer(&__memberName)), __element))
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__className)
			plugify.DestroyString(&__memberName)
		},
	}.Do()
	return __retVal
}

// SetEntSchemaFloat - Sets a float value in an entity's schema.
// @param entityHandle: The handle of the entity from which the value is to be retrieved.
// @param className: The name of the class.
// @param memberName: The name of the schema member.
// @param value: The float value to set.
// @param changeState: If true, change will be sent over the network.
// @param element: Element # (starting from 0) if schema is an array.
func SetEntSchemaFloat(entityHandle int32, className string, memberName string, value float64, changeState bool, element int32) {
	__entityHandle := C.int32_t(entityHandle)
	__className := plugify.ConstructString(className)
	__memberName := plugify.ConstructString(memberName)
	__value := C.double(value)
	__changeState := C.bool(changeState)
	__element := C.int32_t(element)
	plugify.Block{
		Try: func() {
			C.SetEntSchemaFloat(__entityHandle, (*C.String)(unsafe.Pointer(&__className)), (*C.String)(unsafe.Pointer(&__memberName)), __value, __changeState, __element)
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__className)
			plugify.DestroyString(&__memberName)
		},
	}.Do()
}

// GetEntSchemaString - Retrieves a string value from an entity's schema.
// @param entityHandle: The handle of the entity from which the value is to be retrieved.
// @param className: The name of the class.
// @param memberName: The name of the schema member.
// @param element: Element # (starting from 0) if schema is an array.
// @return A string value at the given schema offset.
func GetEntSchemaString(entityHandle int32, className string, memberName string, element int32) string {
	var __retVal string
	var __retVal_native plugify.PlgString
	__entityHandle := C.int32_t(entityHandle)
	__className := plugify.ConstructString(className)
	__memberName := plugify.ConstructString(memberName)
	__element := C.int32_t(element)
	plugify.Block{
		Try: func() {
			__native := C.GetEntSchemaString(__entityHandle, (*C.String)(unsafe.Pointer(&__className)), (*C.String)(unsafe.Pointer(&__memberName)), __element)
			__retVal_native = *(*plugify.PlgString)(unsafe.Pointer(&__native))
			// Unmarshal - Convert native data to managed data.
			__retVal = plugify.GetStringData(&__retVal_native)
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__retVal_native)
			plugify.DestroyString(&__className)
			plugify.DestroyString(&__memberName)
		},
	}.Do()
	return __retVal
}

// SetEntSchemaString - Sets a string value in an entity's schema.
// @param entityHandle: The handle of the entity from which the value is to be retrieved.
// @param className: The name of the class.
// @param memberName: The name of the schema member.
// @param value: The string value to set.
// @param changeState: If true, change will be sent over the network.
// @param element: Element # (starting from 0) if schema is an array.
func SetEntSchemaString(entityHandle int32, className string, memberName string, value string, changeState bool, element int32) {
	__entityHandle := C.int32_t(entityHandle)
	__className := plugify.ConstructString(className)
	__memberName := plugify.ConstructString(memberName)
	__value := plugify.ConstructString(value)
	__changeState := C.bool(changeState)
	__element := C.int32_t(element)
	plugify.Block{
		Try: func() {
			C.SetEntSchemaString(__entityHandle, (*C.String)(unsafe.Pointer(&__className)), (*C.String)(unsafe.Pointer(&__memberName)), (*C.String)(unsafe.Pointer(&__value)), __changeState, __element)
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__className)
			plugify.DestroyString(&__memberName)
			plugify.DestroyString(&__value)
		},
	}.Do()
}

// GetEntSchemaVector3D - Retrieves a vector value from an entity's schema.
// @param entityHandle: The handle of the entity from which the value is to be retrieved.
// @param className: The name of the class.
// @param memberName: The name of the schema member.
// @param element: Element # (starting from 0) if schema is an array.
// @return A string value at the given schema offset.
func GetEntSchemaVector3D(entityHandle int32, className string, memberName string, element int32) plugify.Vector3 {
	var __retVal plugify.Vector3
	__entityHandle := C.int32_t(entityHandle)
	__className := plugify.ConstructString(className)
	__memberName := plugify.ConstructString(memberName)
	__element := C.int32_t(element)
	plugify.Block{
		Try: func() {
			__native := C.GetEntSchemaVector3D(__entityHandle, (*C.String)(unsafe.Pointer(&__className)), (*C.String)(unsafe.Pointer(&__memberName)), __element)
			__retVal = *(*plugify.Vector3)(unsafe.Pointer(&__native))
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__className)
			plugify.DestroyString(&__memberName)
		},
	}.Do()
	return __retVal
}

// SetEntSchemaVector3D - Sets a vector value in an entity's schema.
// @param entityHandle: The handle of the entity from which the value is to be retrieved.
// @param className: The name of the class.
// @param memberName: The name of the schema member.
// @param value: The vector value to set.
// @param changeState: If true, change will be sent over the network.
// @param element: Element # (starting from 0) if schema is an array.
func SetEntSchemaVector3D(entityHandle int32, className string, memberName string, value plugify.Vector3, changeState bool, element int32) {
	__entityHandle := C.int32_t(entityHandle)
	__className := plugify.ConstructString(className)
	__memberName := plugify.ConstructString(memberName)
	__value := *(*C.Vector3)(unsafe.Pointer(&value))
	__changeState := C.bool(changeState)
	__element := C.int32_t(element)
	plugify.Block{
		Try: func() {
			C.SetEntSchemaVector3D(__entityHandle, (*C.String)(unsafe.Pointer(&__className)), (*C.String)(unsafe.Pointer(&__memberName)), &__value, __changeState, __element)
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__className)
			plugify.DestroyString(&__memberName)
		},
	}.Do()
}

// GetEntSchemaVector2D - Retrieves a vector value from an entity's schema.
// @param entityHandle: The handle of the entity from which the value is to be retrieved.
// @param className: The name of the class.
// @param memberName: The name of the schema member.
// @param element: Element # (starting from 0) if schema is an array.
// @return A string value at the given schema offset.
func GetEntSchemaVector2D(entityHandle int32, className string, memberName string, element int32) plugify.Vector2 {
	var __retVal plugify.Vector2
	__entityHandle := C.int32_t(entityHandle)
	__className := plugify.ConstructString(className)
	__memberName := plugify.ConstructString(memberName)
	__element := C.int32_t(element)
	plugify.Block{
		Try: func() {
			__native := C.GetEntSchemaVector2D(__entityHandle, (*C.String)(unsafe.Pointer(&__className)), (*C.String)(unsafe.Pointer(&__memberName)), __element)
			__retVal = *(*plugify.Vector2)(unsafe.Pointer(&__native))
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__className)
			plugify.DestroyString(&__memberName)
		},
	}.Do()
	return __retVal
}

// SetEntSchemaVector2D - Sets a vector value in an entity's schema.
// @param entityHandle: The handle of the entity from which the value is to be retrieved.
// @param className: The name of the class.
// @param memberName: The name of the schema member.
// @param value: The vector value to set.
// @param changeState: If true, change will be sent over the network.
// @param element: Element # (starting from 0) if schema is an array.
func SetEntSchemaVector2D(entityHandle int32, className string, memberName string, value plugify.Vector2, changeState bool, element int32) {
	__entityHandle := C.int32_t(entityHandle)
	__className := plugify.ConstructString(className)
	__memberName := plugify.ConstructString(memberName)
	__value := *(*C.Vector2)(unsafe.Pointer(&value))
	__changeState := C.bool(changeState)
	__element := C.int32_t(element)
	plugify.Block{
		Try: func() {
			C.SetEntSchemaVector2D(__entityHandle, (*C.String)(unsafe.Pointer(&__className)), (*C.String)(unsafe.Pointer(&__memberName)), &__value, __changeState, __element)
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__className)
			plugify.DestroyString(&__memberName)
		},
	}.Do()
}

// GetEntSchemaVector4D - Retrieves a vector value from an entity's schema.
// @param entityHandle: The handle of the entity from which the value is to be retrieved.
// @param className: The name of the class.
// @param memberName: The name of the schema member.
// @param element: Element # (starting from 0) if schema is an array.
// @return A string value at the given schema offset.
func GetEntSchemaVector4D(entityHandle int32, className string, memberName string, element int32) plugify.Vector4 {
	var __retVal plugify.Vector4
	__entityHandle := C.int32_t(entityHandle)
	__className := plugify.ConstructString(className)
	__memberName := plugify.ConstructString(memberName)
	__element := C.int32_t(element)
	plugify.Block{
		Try: func() {
			__native := C.GetEntSchemaVector4D(__entityHandle, (*C.String)(unsafe.Pointer(&__className)), (*C.String)(unsafe.Pointer(&__memberName)), __element)
			__retVal = *(*plugify.Vector4)(unsafe.Pointer(&__native))
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__className)
			plugify.DestroyString(&__memberName)
		},
	}.Do()
	return __retVal
}

// SetEntSchemaVector4D - Sets a vector value in an entity's schema.
// @param entityHandle: The handle of the entity from which the value is to be retrieved.
// @param className: The name of the class.
// @param memberName: The name of the schema member.
// @param value: The vector value to set.
// @param changeState: If true, change will be sent over the network.
// @param element: Element # (starting from 0) if schema is an array.
func SetEntSchemaVector4D(entityHandle int32, className string, memberName string, value plugify.Vector4, changeState bool, element int32) {
	__entityHandle := C.int32_t(entityHandle)
	__className := plugify.ConstructString(className)
	__memberName := plugify.ConstructString(memberName)
	__value := *(*C.Vector4)(unsafe.Pointer(&value))
	__changeState := C.bool(changeState)
	__element := C.int32_t(element)
	plugify.Block{
		Try: func() {
			C.SetEntSchemaVector4D(__entityHandle, (*C.String)(unsafe.Pointer(&__className)), (*C.String)(unsafe.Pointer(&__memberName)), &__value, __changeState, __element)
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__className)
			plugify.DestroyString(&__memberName)
		},
	}.Do()
}

// GetEntSchemaEnt - Retrieves an entity handle from an entity's schema.
// @param entityHandle: The handle of the entity from which the value is to be retrieved.
// @param className: The name of the class.
// @param memberName: The name of the schema member.
// @param element: Element # (starting from 0) if schema is an array.
// @return A string value at the given schema offset.
func GetEntSchemaEnt(entityHandle int32, className string, memberName string, element int32) int32 {
	var __retVal int32
	__entityHandle := C.int32_t(entityHandle)
	__className := plugify.ConstructString(className)
	__memberName := plugify.ConstructString(memberName)
	__element := C.int32_t(element)
	plugify.Block{
		Try: func() {
			__retVal = int32(C.GetEntSchemaEnt(__entityHandle, (*C.String)(unsafe.Pointer(&__className)), (*C.String)(unsafe.Pointer(&__memberName)), __element))
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__className)
			plugify.DestroyString(&__memberName)
		},
	}.Do()
	return __retVal
}

// SetEntSchemaEnt - Sets an entity handle in an entity's schema.
// @param entityHandle: The handle of the entity from which the value is to be retrieved.
// @param className: The name of the class.
// @param memberName: The name of the schema member.
// @param value: The entity handle to set.
// @param changeState: If true, change will be sent over the network.
// @param element: Element # (starting from 0) if schema is an array.
func SetEntSchemaEnt(entityHandle int32, className string, memberName string, value int32, changeState bool, element int32) {
	__entityHandle := C.int32_t(entityHandle)
	__className := plugify.ConstructString(className)
	__memberName := plugify.ConstructString(memberName)
	__value := C.int32_t(value)
	__changeState := C.bool(changeState)
	__element := C.int32_t(element)
	plugify.Block{
		Try: func() {
			C.SetEntSchemaEnt(__entityHandle, (*C.String)(unsafe.Pointer(&__className)), (*C.String)(unsafe.Pointer(&__memberName)), __value, __changeState, __element)
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__className)
			plugify.DestroyString(&__memberName)
		},
	}.Do()
}

// NetworkStateChanged - Updates the networked state of a schema field for a given entity handle.
// @param entityHandle: The handle of the entity from which the value is to be retrieved.
// @param className: The name of the class that contains the member.
// @param memberName: The name of the member to be set.
func NetworkStateChanged(entityHandle int32, className string, memberName string) {
	__entityHandle := C.int32_t(entityHandle)
	__className := plugify.ConstructString(className)
	__memberName := plugify.ConstructString(memberName)
	plugify.Block{
		Try: func() {
			C.NetworkStateChanged(__entityHandle, (*C.String)(unsafe.Pointer(&__className)), (*C.String)(unsafe.Pointer(&__memberName)))
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__className)
			plugify.DestroyString(&__memberName)
		},
	}.Do()
}

// CreateTimer - Creates a new timer that executes a callback function at specified intervals.
// @param interval: The time interval in seconds between each callback execution.
// @param callback: The function to be called when the timer expires.
// @param flags: Flags that modify the behavior of the timer (e.g., no-map change, repeating).
// @param userData: An array intended to hold user-related data, allowing for elements of any type.
// @return A handle to the newly created CTimer object, or -1 if the timer could not be created.
func CreateTimer(interval float64, callback TimerCallback, flags TimerFlag, userData []interface{}) uint64 {
	var __retVal uint64
	__interval := C.double(interval)
	__callback := plugify.GetFunctionPointerForDelegate(callback)
	__flags := C.int32_t(flags)
	__userData := plugify.ConstructVectorVariant(userData)
	plugify.Block{
		Try: func() {
			__retVal = uint64(C.CreateTimer(__interval, __callback, __flags, (*C.Vector)(unsafe.Pointer(&__userData))))
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyVectorVariant(&__userData)
		},
	}.Do()
	return __retVal
}

// KillsTimer - Stops and removes an existing timer.
// @param timer: A handle of the CTimer object to be stopped and removed.
func KillsTimer(timer uint64) {
	__timer := C.uint64_t(timer)
	C.KillsTimer(__timer)
}

// GetTickInterval - Returns the number of seconds in between game server ticks.
// @return The tick interval value.
func GetTickInterval() float64 {
	__retVal := float64(C.GetTickInterval())
	return __retVal
}

// GetTickedTime - Returns the simulated game time.
// @return The ticked time value.
func GetTickedTime() float64 {
	__retVal := float64(C.GetTickedTime())
	return __retVal
}

// OnClientConnect_Register - Register callback to event.
// @param callback: Function callback.
func OnClientConnect_Register(callback OnClientConnectCallback) {
	__callback := plugify.GetFunctionPointerForDelegate(callback)
	C.OnClientConnect_Register(__callback)
}

// OnClientConnect_Unregister - Unregister callback to event.
// @param callback: Function callback.
func OnClientConnect_Unregister(callback OnClientConnectCallback) {
	__callback := plugify.GetFunctionPointerForDelegate(callback)
	C.OnClientConnect_Unregister(__callback)
}

// OnClientConnect_Post_Register - Register callback to event.
// @param callback: Function callback.
func OnClientConnect_Post_Register(callback OnClientConnect_PostCallback) {
	__callback := plugify.GetFunctionPointerForDelegate(callback)
	C.OnClientConnect_Post_Register(__callback)
}

// OnClientConnect_Post_Unregister - Unregister callback to event.
// @param callback: Function callback.
func OnClientConnect_Post_Unregister(callback OnClientConnect_PostCallback) {
	__callback := plugify.GetFunctionPointerForDelegate(callback)
	C.OnClientConnect_Post_Unregister(__callback)
}

// OnClientConnected_Register - Register callback to event.
// @param callback: Function callback.
func OnClientConnected_Register(callback OnClientConnectedCallback) {
	__callback := plugify.GetFunctionPointerForDelegate(callback)
	C.OnClientConnected_Register(__callback)
}

// OnClientConnected_Unregister - Unregister callback to event.
// @param callback: Function callback.
func OnClientConnected_Unregister(callback OnClientConnectedCallback) {
	__callback := plugify.GetFunctionPointerForDelegate(callback)
	C.OnClientConnected_Unregister(__callback)
}

// OnClientPutInServer_Register - Register callback to event.
// @param callback: Function callback.
func OnClientPutInServer_Register(callback OnClientPutInServerCallback) {
	__callback := plugify.GetFunctionPointerForDelegate(callback)
	C.OnClientPutInServer_Register(__callback)
}

// OnClientPutInServer_Unregister - Unregister callback to event.
// @param callback: Function callback.
func OnClientPutInServer_Unregister(callback OnClientPutInServerCallback) {
	__callback := plugify.GetFunctionPointerForDelegate(callback)
	C.OnClientPutInServer_Unregister(__callback)
}

// OnClientDisconnect_Register - Register callback to event.
// @param callback: Function callback.
func OnClientDisconnect_Register(callback OnClientDisconnectCallback) {
	__callback := plugify.GetFunctionPointerForDelegate(callback)
	C.OnClientDisconnect_Register(__callback)
}

// OnClientDisconnect_Unregister - Unregister callback to event.
// @param callback: Function callback.
func OnClientDisconnect_Unregister(callback OnClientDisconnectCallback) {
	__callback := plugify.GetFunctionPointerForDelegate(callback)
	C.OnClientDisconnect_Unregister(__callback)
}

// OnClientDisconnect_Post_Register - Register callback to event.
// @param callback: Function callback.
func OnClientDisconnect_Post_Register(callback OnClientDisconnect_PostCallback) {
	__callback := plugify.GetFunctionPointerForDelegate(callback)
	C.OnClientDisconnect_Post_Register(__callback)
}

// OnClientDisconnect_Post_Unregister - Unregister callback to event.
// @param callback: Function callback.
func OnClientDisconnect_Post_Unregister(callback OnClientDisconnect_PostCallback) {
	__callback := plugify.GetFunctionPointerForDelegate(callback)
	C.OnClientDisconnect_Post_Unregister(__callback)
}

// OnClientActive_Register - Register callback to event.
// @param callback: Function callback.
func OnClientActive_Register(callback OnClientActiveCallback) {
	__callback := plugify.GetFunctionPointerForDelegate(callback)
	C.OnClientActive_Register(__callback)
}

// OnClientActive_Unregister - Unregister callback to event.
// @param callback: Function callback.
func OnClientActive_Unregister(callback OnClientActiveCallback) {
	__callback := plugify.GetFunctionPointerForDelegate(callback)
	C.OnClientActive_Unregister(__callback)
}

// OnClientFullyConnect_Register - Register callback to event.
// @param callback: Function callback.
func OnClientFullyConnect_Register(callback OnClientFullyConnectCallback) {
	__callback := plugify.GetFunctionPointerForDelegate(callback)
	C.OnClientFullyConnect_Register(__callback)
}

// OnClientFullyConnect_Unregister - Unregister callback to event.
// @param callback: Function callback.
func OnClientFullyConnect_Unregister(callback OnClientFullyConnectCallback) {
	__callback := plugify.GetFunctionPointerForDelegate(callback)
	C.OnClientFullyConnect_Unregister(__callback)
}

// OnClientSettingsChanged_Register - Register callback to event.
// @param callback: Function callback.
func OnClientSettingsChanged_Register(callback OnClientSettingsChangedCallback) {
	__callback := plugify.GetFunctionPointerForDelegate(callback)
	C.OnClientSettingsChanged_Register(__callback)
}

// OnClientSettingsChanged_Unregister - Unregister callback to event.
// @param callback: Function callback.
func OnClientSettingsChanged_Unregister(callback OnClientSettingsChangedCallback) {
	__callback := plugify.GetFunctionPointerForDelegate(callback)
	C.OnClientSettingsChanged_Unregister(__callback)
}

// OnClientAuthenticated_Register - Register callback to event.
// @param callback: Function callback.
func OnClientAuthenticated_Register(callback OnClientAuthenticatedCallback) {
	__callback := plugify.GetFunctionPointerForDelegate(callback)
	C.OnClientAuthenticated_Register(__callback)
}

// OnClientAuthenticated_Unregister - Unregister callback to event.
// @param callback: Function callback.
func OnClientAuthenticated_Unregister(callback OnClientAuthenticatedCallback) {
	__callback := plugify.GetFunctionPointerForDelegate(callback)
	C.OnClientAuthenticated_Unregister(__callback)
}

// OnLevelInit_Register - Register callback to event.
// @param callback: Function callback.
func OnLevelInit_Register(callback OnLevelInitCallback) {
	__callback := plugify.GetFunctionPointerForDelegate(callback)
	C.OnLevelInit_Register(__callback)
}

// OnLevelInit_Unregister - Unregister callback to event.
// @param callback: Function callback.
func OnLevelInit_Unregister(callback OnLevelInitCallback) {
	__callback := plugify.GetFunctionPointerForDelegate(callback)
	C.OnLevelInit_Unregister(__callback)
}

// OnLevelShutdown_Register - Register callback to event.
// @param callback: Function callback.
func OnLevelShutdown_Register(callback OnLevelShutdownCallback) {
	__callback := plugify.GetFunctionPointerForDelegate(callback)
	C.OnLevelShutdown_Register(__callback)
}

// OnLevelShutdown_Unregister - Unregister callback to event.
// @param callback: Function callback.
func OnLevelShutdown_Unregister(callback OnLevelShutdownCallback) {
	__callback := plugify.GetFunctionPointerForDelegate(callback)
	C.OnLevelShutdown_Unregister(__callback)
}

// OnEntitySpawned_Register - Register callback to event.
// @param callback: Function callback.
func OnEntitySpawned_Register(callback OnEntitySpawnedCallback) {
	__callback := plugify.GetFunctionPointerForDelegate(callback)
	C.OnEntitySpawned_Register(__callback)
}

// OnEntitySpawned_Unregister - Unregister callback to event.
// @param callback: Function callback.
func OnEntitySpawned_Unregister(callback OnEntitySpawnedCallback) {
	__callback := plugify.GetFunctionPointerForDelegate(callback)
	C.OnEntitySpawned_Unregister(__callback)
}

// OnEntityCreated_Register - Register callback to event.
// @param callback: Function callback.
func OnEntityCreated_Register(callback OnEntityCreatedCallback) {
	__callback := plugify.GetFunctionPointerForDelegate(callback)
	C.OnEntityCreated_Register(__callback)
}

// OnEntityCreated_Unregister - Unregister callback to event.
// @param callback: Function callback.
func OnEntityCreated_Unregister(callback OnEntityCreatedCallback) {
	__callback := plugify.GetFunctionPointerForDelegate(callback)
	C.OnEntityCreated_Unregister(__callback)
}

// OnEntityDeleted_Register - Register callback to event.
// @param callback: Function callback.
func OnEntityDeleted_Register(callback OnEntityDeletedCallback) {
	__callback := plugify.GetFunctionPointerForDelegate(callback)
	C.OnEntityDeleted_Register(__callback)
}

// OnEntityDeleted_Unregister - Unregister callback to event.
// @param callback: Function callback.
func OnEntityDeleted_Unregister(callback OnEntityDeletedCallback) {
	__callback := plugify.GetFunctionPointerForDelegate(callback)
	C.OnEntityDeleted_Unregister(__callback)
}

// OnEntityParentChanged_Register - Register callback to event.
// @param callback: Function callback.
func OnEntityParentChanged_Register(callback OnEntityParentChangedCallback) {
	__callback := plugify.GetFunctionPointerForDelegate(callback)
	C.OnEntityParentChanged_Register(__callback)
}

// OnEntityParentChanged_Unregister - Unregister callback to event.
// @param callback: Function callback.
func OnEntityParentChanged_Unregister(callback OnEntityParentChangedCallback) {
	__callback := plugify.GetFunctionPointerForDelegate(callback)
	C.OnEntityParentChanged_Unregister(__callback)
}

// OnServerStartup_Register - Register callback to event.
// @param callback: Function callback.
func OnServerStartup_Register(callback OnServerStartupCallback) {
	__callback := plugify.GetFunctionPointerForDelegate(callback)
	C.OnServerStartup_Register(__callback)
}

// OnServerStartup_Unregister - Unregister callback to event.
// @param callback: Function callback.
func OnServerStartup_Unregister(callback OnServerStartupCallback) {
	__callback := plugify.GetFunctionPointerForDelegate(callback)
	C.OnServerStartup_Unregister(__callback)
}

// OnServerActivate_Register - Register callback to event.
// @param callback: Function callback.
func OnServerActivate_Register(callback OnServerActivateCallback) {
	__callback := plugify.GetFunctionPointerForDelegate(callback)
	C.OnServerActivate_Register(__callback)
}

// OnServerActivate_Unregister - Unregister callback to event.
// @param callback: Function callback.
func OnServerActivate_Unregister(callback OnServerActivateCallback) {
	__callback := plugify.GetFunctionPointerForDelegate(callback)
	C.OnServerActivate_Unregister(__callback)
}

// OnGameFrame_Register - Register callback to event.
// @param callback: Function callback.
func OnGameFrame_Register(callback OnGameFrameCallback) {
	__callback := plugify.GetFunctionPointerForDelegate(callback)
	C.OnGameFrame_Register(__callback)
}

// OnGameFrame_Unregister - Unregister callback to event.
// @param callback: Function callback.
func OnGameFrame_Unregister(callback OnGameFrameCallback) {
	__callback := plugify.GetFunctionPointerForDelegate(callback)
	C.OnGameFrame_Unregister(__callback)
}

// OnUpdateWhenNotInGame_Register - Register callback to event.
// @param callback: Function callback.
func OnUpdateWhenNotInGame_Register(callback OnUpdateWhenNotInGameCallback) {
	__callback := plugify.GetFunctionPointerForDelegate(callback)
	C.OnUpdateWhenNotInGame_Register(__callback)
}

// OnUpdateWhenNotInGame_Unregister - Unregister callback to event.
// @param callback: Function callback.
func OnUpdateWhenNotInGame_Unregister(callback OnUpdateWhenNotInGameCallback) {
	__callback := plugify.GetFunctionPointerForDelegate(callback)
	C.OnUpdateWhenNotInGame_Unregister(__callback)
}

// OnPreWorldUpdate_Register - Register callback to event.
// @param callback: Function callback.
func OnPreWorldUpdate_Register(callback OnPreWorldUpdateCallback) {
	__callback := plugify.GetFunctionPointerForDelegate(callback)
	C.OnPreWorldUpdate_Register(__callback)
}

// OnPreWorldUpdate_Unregister - Unregister callback to event.
// @param callback: Function callback.
func OnPreWorldUpdate_Unregister(callback OnPreWorldUpdateCallback) {
	__callback := plugify.GetFunctionPointerForDelegate(callback)
	C.OnPreWorldUpdate_Unregister(__callback)
}

// GetGameRules - Retrieves the pointer to the current game rules instance.
// @return A pointer to the game rules object.
func GetGameRules() uintptr {
	__retVal := uintptr(C.GetGameRules())
	return __retVal
}

// TerminateRound - Forces the round to end with a specified reason and delay.
// @param delay: Time (in seconds) to delay before the next round starts.
// @param reason: The reason for ending the round, defined by the CSRoundEndReason enum.
func TerminateRound(delay float32, reason CSRoundEndReason) {
	__delay := C.float(delay)
	__reason := C.int32_t(reason)
	C.TerminateRound(__delay, __reason)
}

// GetWeaponVData - Retrieves the weapon VData for a given weapon.
// @param entityHandle: The handle of the entity from which to retrieve the weapon VData.
// @return The handle of the entity from which to retrieve the weapon VData.
func GetWeaponVData(entityHandle int32) uintptr {
	var __retVal uintptr
	__entityHandle := C.int32_t(entityHandle)
	__retVal = uintptr(C.GetWeaponVData(__entityHandle))
	return __retVal
}

// GetWeaponDefIndex - Retrieves the weapon definition index for a given entity handle.
// @param entityHandle: The handle of the entity from which to retrieve the weapon def index.
// @return The weapon definition index as a `uint16_t`, or 0 if the entity handle is invalid.
func GetWeaponDefIndex(entityHandle int32) uint16 {
	var __retVal uint16
	__entityHandle := C.int32_t(entityHandle)
	__retVal = uint16(C.GetWeaponDefIndex(__entityHandle))
	return __retVal
}
