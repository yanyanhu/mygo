package main


type PersonInfo struct {
    ID string
    Name string
    Address string
}

func searchPerson(personMap map[string] PersonInfo, key string) PersonInfo {
    res, ok := personMap[key]
    if ok {
        println(res.ID)
        println(res.Name)
        println(res.Address)
    } else {
	println("Person with key \"%s\" can not be found.", key)
    }

    return res
}

func main() {
    //var personmap map[string] PersonInfo
    personMap := make(map[string] PersonInfo)
    personMap["1"] = PersonInfo{"0001", "John", "NewYork"}
    personMap["2"] = PersonInfo{"0002", "Lee", "BeiJing"}

    searchPerson(personMap, "1")
    delete(personMap, "1")
    searchPerson(personMap, "1")
}

