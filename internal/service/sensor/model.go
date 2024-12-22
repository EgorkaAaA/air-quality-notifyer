package sensor

import (
	"air-quality-notifyer/internal/lib"
	"fmt"
)

type AirqualitySensor struct {
	Id                         int
	Date                       string
	SDS_P2                     float64
	SDS_P1                     float64
	Temperature                float64
	Humidity                   int
	Pressure                   int
	District                   string
	AQIPM25                    float64
	AQIPM10                    float64
	AQIPM10WarningIndex        int64
	AQIPM25WarningIndex        int64
	DangerLevel                string
	DangerColor                string
	AdditionalInfo             string
	AQIPM10Analysis            string
	AQIPM25Analysis            string
	AQIAnalysisRecommendations string
	SourceLink                 string
}

type pmLevelAir struct {
	PM25Low                    float64
	PM25High                   float64
	PM10Low                    float64
	PM10High                   float64
	IndexLow                   float64
	IndexHigh                  float64
	Color                      string
	Name                       string
	AQIAnalysis                string
	AQIAnalysisRecommendations string
}

var pmLevelAirMap = []pmLevelAir{
	{
		PM25Low:                    0,
		PM25High:                   12,
		PM10Low:                    0,
		PM10High:                   54,
		IndexLow:                   0,
		IndexHigh:                  50,
		Color:                      "#50ccaa",
		Name:                       "Хорошо",
		AQIAnalysis:                "Нормальный уровень",
		AQIAnalysisRecommendations: "Отличный день для активного отдыха на свежем воздухе",
	},
	{
		PM25Low:                    12.1,
		PM25High:                   35.4,
		PM10Low:                    55,
		PM10High:                   154,
		IndexLow:                   51,
		IndexHigh:                  100,
		Color:                      "#f0e641",
		Name:                       "Приемлемо",
		AQIAnalysis:                "Нормальный уровень",
		AQIAnalysisRecommendations: "Некоторые люди могут быть чувствительны к загрязнению частицами.\n\n<b>Чувствительные люди</b>: попробуйте уменьшить длительные или тяжелые нагрузки. Следите за такими симптомами, как кашель или одышка. Это признаки того, что нужно снизить нагрузку.\n\n<b>Всем остальным</b>: это хороший день для активности на улице.",
	},
	{
		PM25Low:                    35.5,
		PM25High:                   55.4,
		PM10Low:                    155,
		PM10High:                   254,
		IndexLow:                   101,
		IndexHigh:                  150,
		Color:                      "#fa912a",
		Name:                       "Плохо",
		AQIAnalysis:                "Повышенный уровень - \"плохо\" ⚠️",
		AQIAnalysisRecommendations: "К уязвимым группам относятся люди <b>с заболеваниями сердца или легких, пожилые люди, дети и подростки</b>.\n\n<b>Чувствительные группы</b>: уменьшите длительные или тяжелые нагрузки. Активный образ жизни на улице - это нормально, но делайте больше перерывов и делайте менее интенсивные занятия. Следите за такими симптомами, как кашель или одышка.\n\n<b>Люди, страдающие астмой</b>, должны следовать своим планам действий при астме и иметь под рукой лекарства быстрого действия.\n\n<b>Если у вас заболевание сердца</b>: такие симптомы, как учащенное сердцебиение, одышка или необычная усталость, могут указывать на серьезную проблему. Если у вас есть какие-либо из них, обратитесь к своему врачу.",
	},
	{
		PM25Low:                    55.5,
		PM25High:                   150.4,
		PM10Low:                    255,
		PM10High:                   354,
		IndexLow:                   151,
		IndexHigh:                  200,
		Color:                      "#ff5050",
		Name:                       "Вредно",
		AQIAnalysis:                "Повышенный уровень - \"вредно\" ⚠️",
		AQIAnalysisRecommendations: "<b>Касается всех</b>\n\n<b>Чувствительные группы</b>: Избегайте длительных или тяжелых нагрузок. Подумайте о том, чтобы переместиться в помещение или изменить расписание.\n\n<b>Всем остальным</b>: уменьшите длительные или тяжелые нагрузки. Делайте больше перерывов во время активного отдыха.",
	},
	{
		PM25Low:                    150.5,
		PM25High:                   250.4,
		PM10Low:                    355,
		PM10High:                   424,
		IndexLow:                   201,
		IndexHigh:                  300,
		Color:                      "#8f3f97",
		Name:                       "Очень вредно",
		AQIAnalysis:                "Опасный уровень - \"очень вредно\" 💀",
		AQIAnalysisRecommendations: "<b>Касается всех</b>\n\n<b>Чувствительные группы</b>: избегайте любых физических нагрузок на открытом воздухе. Переместите занятия в закрытое помещение или перенесите время, когда качество воздуха будет лучше.\n\n<b>Всем остальным</b>: Избегайте длительных или тяжелых нагрузок. Подумайте о том, чтобы переместиться в помещение или перенести время на то время, когда качество воздуха будет лучше.",
	},
	{
		PM25Low:                    250.5,
		PM25High:                   350.4,
		PM10Low:                    425,
		PM10High:                   504,
		IndexLow:                   301,
		IndexHigh:                  400,
		Color:                      "#960032",
		Name:                       "Чрезвычайно опасно",
		AQIAnalysis:                "Опасный уровень - \"чрезвычайно опасно\" 💀💀💀",
		AQIAnalysisRecommendations: "<b>Для всех</b>: избегайте любых физических нагрузок на открытом воздухе.\n\n<b>Чувствительные группы</b>: оставайтесь в помещении и сохраняйте низкий уровень активности. Следуйте советам по сохранению низкого уровня частиц в помещении.",
	},
	{
		PM25Low:                    350.5,
		PM25High:                   500.4,
		PM10Low:                    505,
		PM10High:                   604,
		IndexLow:                   401,
		IndexHigh:                  500,
		Color:                      "#960032",
		Name:                       "Опасно для жизни",
		AQIAnalysis:                "Опасный для жизни уровень 😵",
		AQIAnalysisRecommendations: "<b>Для всех</b>: избегайте любых физических нагрузок на открытом воздухе.\n\n<b>Чувствительные группы</b>: оставайтесь в помещении и сохраняйте низкий уровень активности. Следуйте советам по сохранению низкого уровня частиц в помещении.",
	},
}

func (s *AirqualitySensor) withDistrict(districtName string) {
	s.District = districtName
}

func (s *AirqualitySensor) withApiData(id int64) {
	s.Id = int(id)
	s.SourceLink = fmt.Sprintf("https://airkemerovo.ru/sensor/%d", id)
	if s.Humidity >= 90 {
		s.AdditionalInfo = "Высокая влажность. Показания PM могут быть не корректны\n"
	}
	if s.Temperature < -60 {
		s.AdditionalInfo += fmt.Sprintf("Датчики температуры в районе %s не исправен!\n", s.District)
	}

	for index, pm := range pmLevelAirMap {
		isPM10Dangerous := s.SDS_P1 >= pm.PM10Low && s.SDS_P1 < pm.PM10High
		isPM25Dangerous := s.SDS_P2 >= pm.PM25Low && s.SDS_P2 < pm.PM25High

		if isPM10Dangerous || isPM25Dangerous {
			s.DangerLevel = pm.Name
			s.DangerColor = pm.Color
			s.AQIAnalysisRecommendations = pm.AQIAnalysisRecommendations
		}

		if isPM10Dangerous {
			s.AQIPM10 = lib.CalcAQI(s.SDS_P1, pm.PM10High, pm.PM10Low, pm.IndexHigh, pm.IndexLow)
			s.AQIPM10Analysis = pm.AQIAnalysis
			s.AQIPM10WarningIndex = int64(index)
		}

		if isPM25Dangerous {
			s.AQIPM25 = lib.CalcAQI(s.SDS_P2, pm.PM25High, pm.PM25Low, pm.IndexHigh, pm.IndexLow)
			s.AQIPM25Analysis = pm.AQIAnalysis
			s.AQIPM25WarningIndex = int64(index)
		}
	}

	if s.AQIPM10WarningIndex >= s.AQIPM25WarningIndex {
		s.AQIAnalysisRecommendations = pmLevelAirMap[s.AQIPM10WarningIndex].AQIAnalysisRecommendations
	} else {
		s.AQIAnalysisRecommendations = pmLevelAirMap[s.AQIPM25WarningIndex].AQIAnalysisRecommendations
	}
}
