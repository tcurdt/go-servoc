package main

import ()

type Write struct {
	Address uint16
	Name    string
	Value   uint16
}

func control_mode(m Control_mode) (uint16, uint16) {
	switch m {
	case pos_pulse_ccw_dir:
		return 0x0000, 0x0000
	case pos_pulse_cw_dir:
		return 0x0000, 0x0002
	case pos_dipulse_ccw_dir:
		return 0x0000, 0x0001
	case pos_dipulse_cw_dir:
		return 0x0000, 0x0003
	case pos_digital:
		return 0x0020, 0x0000
	case pos_ramp_enable:
		return 0x0060, 0x0000
	case vel_analog_cw:
		return 0x000c, 0x0000
	case vel_analog_ccw:
		return 0x000c, 0x0004
	case vel_digital_pwm_half:
		return 0x0004, 0x0008
	case vel_digital_pwm_inversion:
		return 0x0004, 0x0010
	case vel_ramp_enable:
		return 0x0084, 0x0010
	case torque_analog_cw:
		return 0x0003, 0x0000
	case torque_analog_ccw:
		return 0x0003, 0x0004
	case torque_digital_pwm_half:
		return 0x0001, 0x0008
	case torque_digital_pwm_inversion:
		return 0x0001, 0x0010
	default:
		panic("unrecognized control mode - please report")
	}
}

func mode2(m struct {
	Control    Mode2_control
	Pos_filter bool
	Vel_filter bool
}) uint16 {
	val := uint16(0)
	if m.Control == external {
		val += 1 << 0
	}
	if m.Pos_filter {
		val += 1 << 1
	}
	if m.Vel_filter {
		val += 1 << 2
	}
	return val
}

func (config Config) Writes() []Write {

	control_mode_6, control_mode_7 := control_mode(config.Control_mode)
	mode2 := mode2(config.Mode2)

	return []Write{
		{0x000a, "motor_encoder.encoder", config.Motor_encoder.Line},
		{0x0046, "eletronic_gear.molecular", config.Eletronic_gear.Molecular},
		{0x0047, "eletronic_gear.denominator", config.Eletronic_gear.Denominator},
		{0x0031, "input_offset", config.Input_offset},
		{0x0032, "weight_coefficient", config.Weight_coefficient},
		{0x0006, "control_mode 1/2", control_mode_6},
		{0x0007, "control_mode 2/2", control_mode_7},
		{0x0008, "mode2", mode2},
		{0x0040, "pos_loop.pp", config.Pos_loop.Pp},
		{0x0041, "pos_loop.pd", config.Pos_loop.Pd},
		{0x0042, "pos_loop.pf", config.Pos_loop.Pf},
		{0x0045, "pos_loop.pos_filter", config.Pos_loop.Pos_filter},
		{0x0048, "pos_loop.pos_error", config.Pos_loop.Pos_error},
		{0x0050, "vel_loop.vp", config.Vel_loop.Vp},
		{0x0055, "vel_loop.continous_vp", config.Vel_loop.Continous_vp},
		{0x0051, "vel_loop.vi", config.Vel_loop.Vi},
		{0x0056, "vel_loop.vel_limit", config.Vel_loop.Vel_limit},
		{0x0052, "vel_loop.vd", config.Vel_loop.Vd},
		{0x0057, "vel_loop.acc", config.Vel_loop.Acc},
		{0x0053, "vel_loop.aff", config.Vel_loop.Aff},
		{0x0058, "vel_loop.dec", config.Vel_loop.Dec},
		{0x0054, "vel_loop.vel_filter", config.Vel_loop.Vel_filter},
		{0x0060, "current_loop.cp", config.Current_loop.Cp},
		{0x0062, "current_loop.continous_current", config.Current_loop.Continous_current},
		{0x0061, "current_loop.ci", config.Current_loop.Ci},
		{0x0063, "current_loop.limit_current", config.Current_loop.Limit_current},
		{0x003a, "threshold_setting.temp_limit", config.Threshold_setting.Temp_limit},
		{0x003b, "threshold_setting.over_voltage_limit", config.Threshold_setting.Over_voltage_limit},
		{0x003c, "threshold_setting.under_voltage_limit", config.Threshold_setting.Under_voltage_limit},
		{0x003d, "threshold_setting.i2t_limit", config.Threshold_setting.I2t_limit},
	}
}
