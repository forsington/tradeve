package item

import (
	"strings"

	"github.com/forsington/tradeve/config"
)

type Type struct {
	TypeId   int    `csv:"typeID"`
	TypeName string `csv:"typeName"`

	GroupId int `csv:"groupID"`
}

type Types []*Type

func (t Types) Active(hasActiveOrders []int32) Types {
	activeTypes := make(Types, 0)

	for _, t := range t {
		for _, h := range hasActiveOrders {
			if t.TypeId == int(h) {
				activeTypes = append(activeTypes, t)
				break
			}
		}
	}

	return activeTypes
}

func (t Types) ExcludeGroups(exclusions *config.ExcludeGroups) Types {
	excludedTypes := make(Types, 0)

	groupsToExclude := groupIdsFromConfig(exclusions)
	for _, t := range t {
		exclude := false
		for _, g := range groupsToExclude {
			if t.GroupId == g {
				exclude = true
				break
			}
		}

		if exclusions.Deprecated && strings.Contains(t.TypeName, "Deprecated") {
			exclude = true
		}

		if !exclude {
			excludedTypes = append(excludedTypes, t)
		}
	}

	return excludedTypes
}

func groupIdsFromConfig(exclusions *config.ExcludeGroups) []int {
	groups := make([]int, 0)

	if exclusions.Skins {
		groups = append(groups, GroupSkins...)
	}
	if exclusions.Wearables {
		groups = append(groups, GroupWearables...)
	}
	if exclusions.Skillbooks {
		groups = append(groups, GroupSkillbooks...)
	}
	if exclusions.Blueprints {
		groups = append(groups, GroupBlueprints...)
	}
	if exclusions.Skinr {
		groups = append(groups, GroupSkinr...)
	}
	if exclusions.Crates {
		groups = append(groups, GroupCrates...)
	}

	return groups
}

var (
	GroupSkins      = []int{1950, 1951, 1952, 1953, 1954, 1955, 4040}
	GroupWearables  = []int{314, 1091, 1088, 1090, 1089, 53, 72, 205, 68, 71, 62, 98, 326, 328, 767, 766, 769, 76, 43, 517, 1271, 1092, 1670, 1083, 4057, 1084}
	GroupSkillbooks = []int{1210, 266, 273, 272, 1216, 258, 255, 256, 257, 275, 1220, 1241, 268, 1218, 269, 1217, 270, 4734, 278, 1545, 1240, 1213, 303}
	GroupBlueprints = []int{104, 105, 106, 107, 134, 136, 166, 135, 133, 165, 111, 108, 110, 118, 119, 120, 126, 127, 128, 129, 132, 139, 142, 143, 145, 147, 148, 151, 154, 156, 157, 167, 168, 169, 141, 123, 163, 162, 160, 161, 176, 174, 177, 158, 137, 152, 296, 140, 1045, 857, 912, 858, 1048, 841, 856, 860, 855, 854, 891, 871, 853, 718, 1145, 1013, 350, 121, 917, 914, 870, 723, 352, 918, 532, 516, 489, 1137, 1141, 787, 1191, 371, 1157, 1152, 1151, 1155, 724, 348, 349, 360, 218, 342, 343, 344, 345, 346, 223, 224, 131, 130, 347, 356, 401, 400, 726, 725, 722, 504, 503, 487, 490, 477, 478, 27, 537, 525, 915, 535, 1139, 1142, 1144, 486, 1146, 1143, 408, 1147, 643, 651, 1222, 1123, 1162, 1703, 859, 172, 944, 888, 890, 945, 178, 996, 965, 1891, 170, 1160, 1190, 1197, 1200, 1239, 1224, 1227, 1267, 1269, 1268, 1270, 1277, 1295, 1293, 1294, 1305, 1318, 1309, 1317, 1397, 1399, 1462, 1709, 1707, 1708, 1542, 1543, 1679, 1718, 1723, 1810, 1812, 1814, 1209, 1880, 973, 1888, 1948, 1990, 1992, 1993, 1994, 2010, 2019, 2023, 4052, 4064, 4065, 4069, 4066, 4095, 4108, 4118, 4141, 4175, 4188, 314, 351121, 351844, 350858, 35121}
	GroupSkinr      = []int{4725, 4726}
	GroupCrates     = []int{1194, 314, 1818}
)
