package main

import "fmt"

type Vision string

const (
	Camera     Vision = "camera"
	Infrared          = "infrared"
	Ultrasonic        = "ultrasonic"
)

type Leg string

const (
	Wheel   Leg = "wheel"
	Bipedal     = "bipedal"
)

type Weapon string

const (
	Handgun Weapon = "handgun"
	Sniper         = "sniper"
	AK47           = "ak47"
)

type RobotBuilder interface {
	AttachVision(Vision) RobotBuilder
	AttachLeg(Leg) RobotBuilder
	UseWeapon(Weapon) RobotBuilder
	Build() Robot
}

type Robot interface {
	Move() string
	Fire() string
}

type robotBuilder struct {
	vision Vision
	leg    Leg
	weapon Weapon
}

func (rb *robotBuilder) AttachVision(vision Vision) RobotBuilder {
	rb.vision = vision
	return rb
}

func (rb *robotBuilder) AttachLeg(leg Leg) RobotBuilder {
	rb.leg = leg
	return rb
}

func (rb *robotBuilder) UseWeapon(weapon Weapon) RobotBuilder {
	rb.weapon = weapon
	return rb
}

func (rb *robotBuilder) Build() Robot {
	return &robot{
		vision: rb.vision,
		leg:    rb.leg,
		weapon: rb.weapon,
	}
}

type robot struct {
	vision Vision
	leg    Leg
	weapon Weapon
}

func (r *robot) Move() string {
	return "Moving forward on " + string(r.leg)
}

func (r *robot) Fire() string {
	return "Firing " + string(r.weapon) + " at enemy detected by " + string(r.vision)
}

func main() {
	builder := &robotBuilder{}
	robot := builder.AttachVision(Infrared).AttachLeg(Wheel).UseWeapon(AK47).Build()
	fmt.Println(robot.Move())
	fmt.Println(robot.Fire())
}
