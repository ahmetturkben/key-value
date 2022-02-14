# In memory key-value çalışan bir REST-API servisi


# Özellikler
+ key ’i set etmek için bir endpoint
+ key ’i get etmek için bir endpoint
+ Komple data’yı flush etmek için bir endpoint
+ Belirli bir interval’da (N dakikada bir) dosyaya kaydetmeli
+ Uygulama durup tekrar ayağa kalktığında, eğer kaydedilmiş dosya varsa, tekrar varolan verileri hafızaya yükelemeli ( /data.txt)

# Endpointler
+ /api/memory/all/flush [GET]
+ /api/memory/{key} [GET]
+ /api/memory [POST]

# Kullanım

+ Get
    + İstenilen keyi getirmek için
    + curl -X 'GET' \
  'http://localhost:80/api/memory/ahmet' \
  -H 'accept: application/json'
  
+ Add
    + Key eklemek için
    + curl -X 'POST' \
  'http://localhost:80/api/memory' \
  -H 'accept: application/json' \
  -H 'Content-Type: application/json' \
  -d '{
  "key": "ahmet",
  "value": "string"
}'

+ Flush
    + Tüm keyleri temizlemek için
    + curl -X 'GET' \
  'http://localhost:80/api/memory/all/flush' \
  -H 'accept: application/json'

# Çalıştırma

+ Standart
   + go run main.go

