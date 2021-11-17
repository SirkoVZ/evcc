package openwb

import "time"

// predefined openWB topic names
const (
	Timeout           = 15 * time.Second
	HeartbeatInterval = 10 * time.Second // loadpoint only client heartbeat

	// root topic
	RootTopic = "openWB"

	// alive
	TimestampTopic = "Timestamp"

	// status
	PluggedTopic    = "boolPlugStat"
	ChargingTopic   = "boolChargeStat"
	ConfiguredTopic = "boolChargePointConfigured"

	// loadpoint only topics
	HeartbeatTopic     = "heartbeat"
	ChargeCurrentTopic = "Current"
	PhasesTopic        = "U1p3p"
	RfidTopic          = "rfid"

	// charge power
	ChargePowerTopic       = "W"
	ChargeTotalEnergyTopic = "kWhCounter"

	// vehicle
	VehicleSoCTopic = "Soc"

	// general measurements
	PowerTopic   = "W"
	SoCTopic     = "%Soc"
	CurrentTopic = "APhase" // 1..3

	// configuration
	PvConfigured      = "boolPVConfigured"
	BatteryConfigured = "boolHouseBatteryConfigured"
)
