db.dropDatabase()

var countriesEn = [
    {
        "country_name": "JAPAN",
        "schengen_flag": false
    },
    {
        "country_name": "AUSTRALIA",
        "schengen_flag": false
    },
    {
        "country_name": "GERMANY",
        "schengen_flag": true
    }
]
for (var index = 0; index < countriesEn.length; index++) {
    db.getCollection('countries_en').insert(countriesEn[index])
}
