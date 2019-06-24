						var tokenString string
						var jres CommonResponse
						jres, tokenString, err = saveAdminToken(&adminNew)
						if err == nil {
							u.Data["json"] = LoginResponse{Token: tokenString, Name: adminNew.Name, ResponseCode: 200, ResponseDescription: "Registration successful"}
						} else {
							debugMessage("RegisterAdmin Error 1: ", err)
							u.Data["json"] = jres
						}

	/*if onlyMine == "YES" {
		cond1 = cond.And("id__in", icountry_ids...)
	} else {
		cond1 = cond.AndNot("id__in", icountry_ids...)//configs.GLOBAL_CODE)
	}*/