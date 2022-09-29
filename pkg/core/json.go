package core

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

//
//
func mapStrFlat(mapJson *map[string]interface{}, key string) map[string]string{
	mapStr := make(map[string]string);
	for k, v := range *mapJson {
		v_str, ok := v.(string)
		if ok{
			if key != ""{
				mapStr[key + "." + k] = v_str
			}else{
				mapStr[k] = v_str
			}
			continue
		}

		v_map, ok := v.(map[string]interface{})
		if !ok{
			continue;
		}
		mapTmp := mapStrFlat(&v_map, k)
		for k2, v2 := range mapTmp {
			mapStr[k2] = v2
		}
	}
	return mapStr
}

func Json2Map(file_name string) map[string]string {
	log.Println("Json2Map=======>",file_name)
	fc, err := ioutil.ReadFile(file_name);
	if err != nil{
		log.Fatalln("ReadFile error: ", err)
	}
	var mapJson map[string]interface{};
	if err := json.Unmarshal(fc,&mapJson); err != nil{
		log.Fatalln("Convert Map Error: ", err)
	}
	return mapStrFlat(&mapJson, "")
}

func ReverseMap(from *map[string]string) map[string]string{
	newMap := make(map[string]string);

	for k, v := range *from {
		newMap[v] = k;
	}

	return newMap
}