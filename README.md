#GoTestAPI for Learning

Note:
1. Masih menggunkan echo
2. Sudah ready mysql
3. Belom tau cara install package juga setelah clone, silahkan gugling2~
4. Desain, Source Code sangat tidak clean code, tolong nanti development aslinya disesuaikan :D
5. ini pure latian

### POST Request
```http
POST /jadwal HTTP/1.1
Content-Type: application/json
```

### Request Body
```json
{
	"title": "Test Sample",
	"lecturer": "Ustadz Muhammad",
	"start": "2018-09-22T19:42:31+07:00",
	"end": "2018-09-22T19:42:31+07:00",
	"event_date": "2018-09-22T19:42:31+07:00"
}
```

### GET Request
```http
GET /jadwal HTTP/1.1
Content-Type: application/json
```