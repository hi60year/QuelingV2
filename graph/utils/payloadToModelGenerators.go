package utils

import (
	"main/graph/model"
)

type PlatformNotSupportedError struct{}

func (PlatformNotSupportedError) Error() string {
	return "platform info not supported"
}

func PlayerFromPayload(payload *model.PlayerPayload, contestId string) (player *model.Player, err error) {

	extraInfo := payload.ExtraInfo

	if payload.Appendix != nil {
		extraInfo["appendix"] = payload.Appendix
	}

	var platformInfos []*model.PlatformInfo

	for _, info := range payload.PlatformInfos {
		var rankingInfo model.RankingInfo
		if info.RankingInfoType == "MahjongSoul" {
			rankingInfo = model.MahjongSoulRankingInfo{
				Ranking3: info.Ranking3,
				Ranking4: info.Ranking4,
			}
			platformInfos = append(platformInfos, &model.PlatformInfo{
				Platform:    "MahjongSoul",
				Name:        info.Name,
				UID:         info.UID,
				RankingInfo: rankingInfo,
			})
		} else if info.RankingInfoType == "Tenhou" {
			rankingInfo = model.TenhouRankingInfo{
				Ranking3: info.Ranking3,
				Ranking4: info.Ranking4,
			}
			platformInfos = append(platformInfos, &model.PlatformInfo{
				Platform:    "Tenhou",
				Name:        info.Name,
				RankingInfo: rankingInfo,
			})
		} else if info.RankingInfoType == "Tziakcha" {
			platformInfos = append(platformInfos, &model.PlatformInfo{
				Platform: "Tziakcha",
				Name:     info.Name,
			})
		} else {
			return nil, PlatformNotSupportedError{}
		}

	}

	player = &model.Player{
		Name:             payload.Name,
		PlatformInfos:    nil,
		ExtraInfo:        extraInfo,
		ProfessionalCert: nil,
		ContestId:        &contestId,
	}
	return player, err
}
