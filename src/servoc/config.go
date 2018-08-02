package main

//go:generate enumer -type=Control_mode -yaml
type Control_mode uint16

const (
	pos_pulse_ccw_dir Control_mode = iota
	pos_pulse_cw_dir
	pos_dipulse_ccw_dir
	pos_dipulse_cw_dir
	pos_digital
	pos_ramp_enable
	vel_analog_cw
	vel_analog_ccw
	vel_digital_pwm_half
	vel_digital_pwm_inversion
	vel_ramp_enable
	torque_analog_cw
	torque_analog_ccw
	torque_digital_pwm_half
	torque_digital_pwm_inversion
)

//go:generate enumer -type=Mode2_control -yaml
type Mode2_control uint16

const (
	internal Mode2_control = iota
	external
)

type Config struct {
	Motor_encoder struct {
		Line uint16
	}
	Eletronic_gear struct {
		Molecular   uint16
		Denominator uint16
	}
	Input_offset       uint16
	Weight_coefficient uint16
	Control_mode       Control_mode
	Mode2              struct {
		Control    Mode2_control
		Pos_filter bool
		Vel_filter bool
	}
	Pos_loop struct {
		Pp         uint16
		Pd         uint16
		Pf         uint16
		Pos_filter uint16
		Pos_error  uint16
	}
	Vel_loop struct {
		Vp           uint16
		Continous_vp uint16
		Vi           uint16
		Vel_limit    uint16
		Vd           uint16
		Acc          uint16
		Aff          uint16
		Dec          uint16
		Vel_filter   uint16
	}
	Current_loop struct {
		Cp                uint16
		Continous_current uint16
		Ci                uint16
		Limit_current     uint16
	}

	Threshold_setting struct {
		Temp_limit          uint16
		Over_voltage_limit  uint16
		Under_voltage_limit uint16
		I2t_limit           uint16
	}
}
