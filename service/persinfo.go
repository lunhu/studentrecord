package service

import (
	//"fmt"
	"errors"
	"strconv"
	"studentrecord/common"
	"studentrecord/dao"
	"studentrecord/model"

	"github.com/tidwall/gjson"
)
// PersInfoAdd 新增一个人员
func PersInfoAdd(reqJSON *gjson.Result) (cjResponse *common.CjResponse, httpStatusCode int) {

	tx, err := sqlxDB.Beginx()
	if err != nil {
		return common.HandError("500", "")
	}
	var pers_id string
	if persNames, err := checkAndReturnPersNames(reqJSON); err == nil {
	
		if len(persNames) == 0 {
			return common.HandError("400", "")
		}

		for _, persName := range persNames {
			
			persName, err := dao.InsertPersName(tx, persName)
			if err != nil {
				tx.Rollback()
				return common.HandError("500", "")
			}

			pers_id = persName.PersID

			newCjResponse, _ := common.HandItem(*persName, 201)

			if cjResponse == nil {
				
				cjResponse = newCjResponse

			} else {

				cjResponse.Collection.Items = append(cjResponse.Collection.Items, newCjResponse.Collection.Items[0])
			}
			
		}
	} else {
		return common.HandError("400", "")
	}

	if persEmails, err := checkAndReturnPersEmails(reqJSON); err == nil {
		
		for _, persEmail := range persEmails {

			if persEmail.PersID =="" {
				persEmail.PersID = pers_id
			}

			persEmail,err := dao.InsertPersEmail(tx, persEmail)

			if err != nil {
				tx.Rollback()
				
				return common.HandError("500", "")
			}

			newCjResponse, _ := common.HandItem(*persEmail, 201)

			if cjResponse == nil {
			
				cjResponse = newCjResponse

			} else {

				cjResponse.Collection.Items = append(cjResponse.Collection.Items, newCjResponse.Collection.Items[0])
			}

		}

	} else {
		return common.HandError("400", "")
		
	}

	if persPhones, err := checkAndReturnPersPhones(reqJSON); err == nil {
		for _, persPhone := range persPhones {
			
			if persPhone.PersID =="" {
				persPhone.PersID = pers_id
			}

			persPhone,err := dao.InsertPersPhone(tx, persPhone)

			if err != nil {
				tx.Rollback()
				
				return common.HandError("500", "")
			}

			newCjResponse, _ := common.HandItem(*persPhone, 201)

			if cjResponse == nil {
			
				cjResponse = newCjResponse

			} else {

				cjResponse.Collection.Items = append(cjResponse.Collection.Items, newCjResponse.Collection.Items[0])
			}

		}
	} else {
		return common.HandError("400", "")
		
	}

	if persEmploys, err := checkAndReturnPersEmploys(reqJSON); err == nil {
		for _, persEmploy := range persEmploys {
			
			if persEmploy.PersID =="" {
				persEmploy.PersID = pers_id
			}

			if persEmploy.EffSeq == "" {

				eff_seq, err := dao.SelectMaxSeq(sqlxDB, persEmploy.PersID)
		
				if err != nil {
					return common.HandError("500", "")
				} 
			
				persEmploy.EffSeq = strconv.Itoa(int(eff_seq + 1))
			}

			persEmploy,err := dao.InsertPersEmploy(tx, persEmploy)

			if err != nil {
				tx.Rollback()
				
				return common.HandError("500", "")
			}

			newCjResponse, _ := common.HandItem(*persEmploy, 201)

			if cjResponse == nil {
			
				cjResponse = newCjResponse

			} else {

				cjResponse.Collection.Items = append(cjResponse.Collection.Items, newCjResponse.Collection.Items[0])
			}
		}
	} else {
		return common.HandError("400", "")
		
	}
	
	tx.Commit()
	return cjResponse, 201
}

// Check for the existence of name
func checkAndReturnPersNames(reqJSON *gjson.Result) ([]*model.PersName, error) {
	var err error
	if reqJSON.Get(`template.data.#(name=="name")`).Exists() {
		 persNames := []*model.PersName{}
		 var pers_id int64
	   	names := reqJSON.Get(`template.data.#(name=="name")#.prompt`)
	   	names.ForEach(func(key, value gjson.Result) bool {

			persName := &model.PersName{}

			FillPersName(reqJSON.Get(`template.data.#(prompt==`+value.String()+`)#`).Array(), persName)

			if persName.PersID == "" {

				if pers_id == 0 {
					pers_id, err = dao.SelectMaxPersId(sqlxDB)
					
					if err != nil {
						err = errors.New(persName.FaildMessage())
						return false
					} 
			
					pers_id += 1
					
				} 

				persName.PersID = strconv.Itoa(int(pers_id))
			}

		   if persName.Valid() {
				
				persNames = append(persNames, persName)
				
		   } else {
				err = errors.New(persName.FaildMessage())
				return false
		   }
		   return true
	   })
	  
		return persNames, err

   } 

   return nil, nil
   
}

func checkAndReturnPersEmails(reqJSON *gjson.Result) ([]*model.PersEmail, error) {
	var err error
	if reqJSON.Get(`template.data.#(name=="email")`).Exists() {
		var persEmails []*model.PersEmail
		   names := reqJSON.Get(`template.data.#(name=="email")#.prompt`)
	   	names.ForEach(func(key, value gjson.Result) bool {
		   //panic(len(reqJSON.Get(`template.data.#(prompt==`+value.String()+`)#`).Array()))
		   persEmail := &model.PersEmail{}

		   FillPersEmail(reqJSON.Get(`template.data.#(prompt==`+value.String()+`)#`).Array(), persEmail)

		   if persEmail.Valid() {
				persEmails = append(persEmails, persEmail)
				
		   } else {
				err = errors.New(persEmail.FaildMessage())
				return false
		   }

		   return true
	   })
	   
		return persEmails, err

   } 
   return nil, nil
}


func checkAndReturnPersPhones(reqJSON *gjson.Result) ([]*model.PersPhone, error) {
	var err error
	if reqJSON.Get(`template.data.#(name=="phone")`).Exists() {
		var persPhones []*model.PersPhone
		   names := reqJSON.Get(`template.data.#(name=="phone")#.prompt`)
	   	names.ForEach(func(key, value gjson.Result) bool {
		   //panic(len(reqJSON.Get(`template.data.#(prompt==`+value.String()+`)#`).Array()))
		   persPhone := &model.PersPhone{}

		   FillPersPhone(reqJSON.Get(`template.data.#(prompt==`+value.String()+`)#`).Array(), persPhone)

		   if persPhone.Valid() {
				persPhones = append(persPhones, persPhone)
				
		   } else {
				err = errors.New(persPhone.FaildMessage())
				return false
		   }

		   return true
	   })
	   
		return persPhones, err

   } 
   return nil, nil
}

func checkAndReturnPersEmploys(reqJSON *gjson.Result) ([]*model.PersEmploy, error) {
	var err error
	if reqJSON.Get(`template.data.#(name=="employer")`).Exists() {
		var persEmploys []*model.PersEmploy
		   names := reqJSON.Get(`template.data.#(name=="employer")#.prompt`)
	   	names.ForEach(func(key, value gjson.Result) bool {
		   //panic(len(reqJSON.Get(`template.data.#(prompt==`+value.String()+`)#`).Array()))
		   persEmploy := &model.PersEmploy{}

		   FillPersEmploy(reqJSON.Get(`template.data.#(prompt==`+value.String()+`)#`).Array(), persEmploy)
		
		   if persEmploy.Valid() {
				persEmploys = append(persEmploys, persEmploy)
				
		   } else {
				err = errors.New(persEmploy.FaildMessage())
				return false
		   }

		   return true
	   })
	   
		return persEmploys, err

   } 
   return nil, nil
}