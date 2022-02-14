curl -X 'POST' \
  'http://localhost:80/api/records' \
  -H 'accept: application/json' \
  -H 'Content-Type: application/json' \
  -d '{
  "startDate": "2015-11-27",
  "endDate": "2015-11-29",
  "minCount": 2084,
  "maxCount": 2086
}'