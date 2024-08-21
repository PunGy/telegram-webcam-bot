package helpers

import (
	"io/ioutil"
	"log"
	"sync"

	"github.com/imdario/mergo"
	"gopkg.in/yaml.v2"
)

type ConfigT struct {
	Basic struct {
		Host     string `yaml:"host"`
		HostId   int64  `yaml:"host-id"`
		Client   string `yaml:"client"`
		ClientId int64  `yaml:"client-id"`
		BotKey   string `yaml:"bot-key"`
	}

	Keyboard struct {
		General struct {
			RequestPhoto         string `yaml:"request-photo"`
			RequestPhotoSequence string `yaml:"request-photo-sequence"`
			RequestVideo         string `yaml:"request-video"`
			Help                 string `yaml:"help"`
			Home                 string `yaml:"home"`
		}

		Host struct {
			ClientInteractionMenu string `yaml:"client-interaction-menu"`
			LittleLeave           string `yaml:"little-leave"`
			ImBack                string `yaml:"im-back"`
			Lunch                 string `yaml:"lunch"`
			WentHome              string `yaml:"went-home"`
			ShowMyself            string `yaml:"show-myself"`
			ShowVideo             string `yaml:"show-video"`
		}

		Client struct {
			Ping string `yaml:"ping"`
		}
	}

	Webcam struct {
		DeviceID string `yaml:"device-id"`
	}

	Responses struct {
		Help                string `yaml:"help"`
		ChooseVideoDuration string `yaml:"choose-video-duration"`
		ChooseSequenceSize  string `yaml:"choose-sequence-size"`
		NumberExpected      string `yaml:"number-expected"`
		LittleLeave         string `yaml:"little-leave"`
		ImBack              string `yaml:"im-back"`
		Lunch               string `yaml:"lunch"`
		WentHome            string `yaml:"went-home"`
		ShowMyself          string `yaml:"show-myself"`
	}
}

var configOnce sync.Once

var cfg ConfigT

func Config() *ConfigT {
	configOnce.Do(func() {
		root := GetDir("../..")

		defaultYaml, err := ioutil.ReadFile(root + "/config/config.default.yml")
		customYaml, customYmlErr := ioutil.ReadFile(root + "/config/config.custom.yml")
		if err != nil {
			log.Printf("Unable to get default config #%v", err)
		}

		err = yaml.Unmarshal(defaultYaml, &cfg)
		if err != nil {
			log.Fatalf("Cannot parse config: %v", err)
		}

		if customYmlErr == nil {
			customCfg := ConfigT{}
			err = yaml.Unmarshal(customYaml, &customCfg)
			if err != nil {
				log.Fatalf("Cannot parse custom config: %v", err)
				return
			}

			if err = mergo.Merge(&cfg, customCfg, mergo.WithOverride); err != nil {
				log.Printf("Cannot merge configs: %v", err)
			}
		}
	})

	return &cfg
}
