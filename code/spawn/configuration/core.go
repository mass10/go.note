package configuration

import "os"

type ConfigurationSettings interface {
	IsChildMode() bool
}

type configurationSettingsCore struct {
	isChildMode bool
}

var _conf *configurationSettingsCore = nil

func Configure() ConfigurationSettings {
	// 一度しか初期化しない
	if _conf != nil {
		return _conf
	}

	_conf = &configurationSettingsCore{}

	for _, e := range os.Args {
		if e == "--fork" {
			_conf.isChildMode = true
		}
	}

	return _conf
}

func (self *configurationSettingsCore) IsChildMode() bool {
	return self.isChildMode
}
