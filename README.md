# servoc

Command line tool to interact with JMC IHSV57 servos.

Uploading you config is as easy as:

```
servo -c config.yaml -p /dev/cu.serial
```

The full configuration looks like this:

```
motor_encoder:
  line: 0
eletronic_gear:
  molecular: 0
  denominator: 0

input_offset: 0
weight_coefficient: 0

control_mode: pos_pulse_ccw_dir
# control_mode: pos_pulse_cw_dir
# control_mode: pos_dipulse_ccw_dir
# control_mode: pos_dipulse_cw_dir
# control_mode: pos_digital
# control_mode: pos_ramp_enable
# control_mode: vel_analog_cw
# control_mode: vel_analog_ccw
# control_mode: vel_digital_pwm_half
# control_mode: vel_digital_pwm_inversion
# control_mode: vel_ramp_enable
# control_mode: torque_analog_cw
# control_mode: torque_analog_ccw
# control_mode: torque_digital_pwm_half
# control_mode: torque_digital_pwm_inversion

mode2:
  control: external
  # control: internal
  pos_filter: false
  vel_filter: false

pos_loop:
  pp: 0
  pd: 0
  pf: 0
  pos_filter: 0
  pos_error: 0

vel_loop:
  vp: 0
  continous_vp: 0
  vi: 0
  vel_limit: 0
  vd: 0
  acc: 0
  aff: 0
  dec: 0
  vel_filter: 0

current_loop:
  cp: 0
  continous_current: 0
  ci: 0
  limit_current: 0

threshold_setting:
  temp_limit: 0
  over_voltage_limit: 0
  under_voltage_limit: 0
  i2t_limit: 0
```

For a more realistic configuration check out [the config from Sebastian End](https://github.com/tcurdt/go-servoc/blob/master/src/sebastian.yaml).


## Where to ask questions

[via gitter](https://gitter.im/tcurdt/go-servoc)

[via git issue](https://github.com/tcurdt/go-servoc/issues)

## License

All code and data is released under the Apache License 2.0.