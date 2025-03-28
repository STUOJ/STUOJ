package judge0

import (
	"encoding/json"
)

func GetLanguage() ([]JudgeLanguage, error) {
	bodystr, err := httpInteraction("/languages", "GET", nil)
	if err != nil {
		return nil, err
	}
	var languages []JudgeLanguage
	err = json.Unmarshal([]byte(bodystr), &languages)
	if err != nil {
		return nil, err
	}
	return languages, nil
}

func GetConfigInfo() (JudgeConfigInfo, error) {
	bodystr, err := httpInteraction("/config_info", "GET", nil)
	if err != nil {
		return JudgeConfigInfo{}, err
	}
	var config JudgeConfigInfo
	err = json.Unmarshal([]byte(bodystr), &config)
	if err != nil {
		return JudgeConfigInfo{}, err
	}
	return config, nil
}

func GetSystemInfo() (JudgeSystemInfo, error) {
	bodystr, err := httpInteraction("/system_info", "GET", nil)
	if err != nil {
		return JudgeSystemInfo{}, err
	}
	var system JudgeSystemInfo
	err = json.Unmarshal([]byte(bodystr), &system)
	if err != nil {
		return JudgeSystemInfo{}, err
	}
	return system, nil
}

func GetStatistics() (JudgeStatistics, error) {
	bodystr, err := httpInteraction("/statistics", "GET", nil)
	if err != nil {
		return JudgeStatistics{}, err
	}
	var statistics JudgeStatistics
	err = json.Unmarshal([]byte(bodystr), &statistics)
	if err != nil {
		return JudgeStatistics{}, err
	}
	return statistics, nil
}

func GetAbout() (JudgeAbout, error) {
	bodystr, err := httpInteraction("/about", "GET", nil)
	if err != nil {
		return JudgeAbout{}, err
	}
	var about JudgeAbout
	err = json.Unmarshal([]byte(bodystr), &about)
	if err != nil {
		return JudgeAbout{}, err
	}
	return about, nil
}

func GetWorkers() ([]JudgeWorker, error) {
	bodystr, err := httpInteraction("/workers", "GET", nil)
	if err != nil {
		return nil, err
	}
	var workers []JudgeWorker
	err = json.Unmarshal([]byte(bodystr), &workers)
	if err != nil {
		return nil, err
	}
	return workers, nil
}

func GetLicense() (string, error) {
	bodystr, err := httpInteraction("/license", "GET", nil)
	if err != nil {
		return "", err
	}
	return bodystr, nil
}

func GetIsolate() (string, error) {
	bodystr, err := httpInteraction("/isolate", "GET", nil)
	if err != nil {
		return "", err
	}
	return bodystr, nil
}

func GetVersion() (string, error) {
	bodystr, err := httpInteraction("/version", "GET", nil)
	if err != nil {
		return "", err
	}
	return bodystr, nil
}
