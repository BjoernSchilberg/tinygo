// +build sam,atsamd21,gamebuino

package machine

import (
	"device/sam"
	"runtime/interrupt"
)

// used to reset into bootloader
const RESET_MAGIC_VALUE = 0xf01669ef

// GPIO Pins
const (
	D0  = PA11  // UART0 RX
	D1  = PA10  // UART0 TX
	D2  = NoPin // does not seem to exist
	D3  = PA09
	D4  = PA08
	D5  = PA15  // PWM available
	D6  = PA20  // PWM available
	D7  = NoPin // does not seem to exist
	D8  = PA06
	D9  = PA07 // PWM available
	D10 = PA18 // can be used for PWM or UART1 TX
	D11 = PA16 // can be used for PWM or UART1 RX
	D12 = PA19 // PWM available
	D13 = PA17 // PWM available
)

// Analog pins
const (
	A0 = PA02 // ADC/AIN[0]
	A1 = PB08 // ADC/AIN[2]
	A2 = PB09 // ADC/AIN[3]
	A3 = PA04 // ADC/AIN[4]
	A4 = PA05 // ADC/AIN[5]
	A5 = PB02 // ADC/AIN[10]
)

const (
	LED = D13
	NEOPIXELS = D8

	//LIGHTSENSOR = A7

	BUTTON_LATCH = PB00
	BUTTON_OUT   = PB30
	BUTTON_CLK   = PB31

	TFT_DC   = PB05
	TFT_CS   = PB07
	TFT_RST  = PA00
	TFT_LITE = PA01

	SPEAKER_ENABLE = PA27

	QSPI_SCK    = PB10
	QSPI_CS     = PB11
	QSPI_DATA_1 = PA08
	QSPI_DATA_2 = PA09
	QSPI_DATA_3 = PA10
	QSPI_DATA_4 = PA11
)

const (
	BUTTON_LEFT_MASK   = 1
	BUTTON_UP_MASK     = 2
	BUTTON_DOWN_MASK   = 4
	BUTTON_RIGHT_MASK  = 8
	BUTTON_SELECT_MASK = 16
	BUTTON_START_MASK  = 32
	BUTTON_A_MASK      = 64
	BUTTON_B_MASK      = 128
)

// UART0 aka USBCDC pins
const (
	USBCDC_DM_PIN = PA24
	USBCDC_DP_PIN = PA25
)

// UART1 pins
const (
	UART_TX_PIN = D10
	UART_RX_PIN = D11
)

// UART1 on the Feather M0.
var (
	UART1 = UART{
		Buffer: NewRingBuffer(),
		Bus:    sam.SERCOM1_USART,
		SERCOM: 1,
	}
)

func init() {
	UART1.Interrupt = interrupt.New(sam.IRQ_SERCOM1, UART1.handleInterrupt)
}

// I2C pins
const (
	SDA_PIN = PA22 // SDA: SERCOM3/PAD[0]
	SCL_PIN = PA23 // SCL: SERCOM3/PAD[1]
)

// I2C on the Feather M0.
var (
	I2C0 = I2C{
		Bus:    sam.SERCOM3_I2CM,
		SERCOM: 3,
	}
)

// SPI pins
const (
	SPI0_SCK_PIN  = PB11 // SCK: SERCOM4/PAD[3]
	SPI0_MOSI_PIN = PB10 // MOSI: SERCOM4/PAD[2]
	SPI0_MISO_PIN = PA12 // MISO: SERCOM4/PAD[0]
)

// SPI on the Feather M0.
var (
	SPI0 = SPI{
		Bus:    sam.SERCOM4_SPI,
		SERCOM: 4,
	}
)

// I2S pins
const (
	I2S_SCK_PIN = PA10
	I2S_SD_PIN  = PA08
	I2S_WS_PIN  = NoPin // TODO: figure out what this is on Feather M0.
)

