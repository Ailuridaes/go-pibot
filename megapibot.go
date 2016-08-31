package main

import (
	"github.com/hybridgroup/gobot/platforms/megapi"
	"log"
)

var _ Robot = (*MegaPiBot)(nil)

const MAX_MOTOR_SPEED = 250
const MIN_MOTOR_SPEED = 50

type MegaPiBot struct {
	leftMotor  *megapi.MotorDriver
	rightMotor *megapi.MotorDriver
}

func NewMegaPiBot(leftMotor *megapi.MotorDriver, rightMotor *megapi.MotorDriver) Robot {
	return &MegaPiBot{
		leftMotor:  leftMotor,
		rightMotor: rightMotor,
	}
}

func (robot *MegaPiBot) HandleCommand(command RobotMoveCommand) {
	if command.Direction == Right || command.Direction == Left {
		command.Speed /= 4
	}
	adjustedSpeed := command.Speed*(MAX_MOTOR_SPEED-MIN_MOTOR_SPEED)/100 + MIN_MOTOR_SPEED
	switch command.Direction {
	case Forward:
		log.Printf("ROBOT: forward at speed %+v\n", adjustedSpeed)
		robot.leftMotor.Speed(-int16(adjustedSpeed))
		robot.rightMotor.Speed(-int16(adjustedSpeed))
	case Backwards:
		log.Printf("ROBOT: backwards at speed %+v\n", adjustedSpeed)
		robot.leftMotor.Speed(+int16(adjustedSpeed))
		robot.rightMotor.Speed(+int16(adjustedSpeed))
	case Left:
		log.Printf("ROBOT: left at speed %+v\n", adjustedSpeed)
		robot.leftMotor.Speed(+int16(adjustedSpeed))
		robot.rightMotor.Speed(-int16(adjustedSpeed))
	case Right:
		log.Printf("ROBOT: right at speed %+v\n", adjustedSpeed)
		robot.leftMotor.Speed(-int16(adjustedSpeed))
		robot.rightMotor.Speed(+int16(adjustedSpeed))
	case Stop:
		log.Println("ROBOT: stopping")
		robot.leftMotor.Speed(0)
		robot.rightMotor.Speed(0)
	case ForwardRight:
		log.Printf("ROBOT: forward-right at speed %+v\n", adjustedSpeed)
	case ForwardLeft:
		log.Printf("ROBOT: forward-left at speed %+v\n", adjustedSpeed)
	case BackwardsRight:
		log.Printf("ROBOT: backwards-right at speed %+v\n", adjustedSpeed)
	case BackwardsLeft:
		log.Printf("ROBOT: backwards-left at speed %+v\n", adjustedSpeed)
	default:
		log.Println("ROBOT: unknown move command")
	}
}
