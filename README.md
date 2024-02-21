# plates.go
The plates.go program is designed to determine the combination of weight plates required to set a barbell to a specific target weight.

### usage
```sh
go run plates/main.go plates/plates.go -w 195
go run plates/main.go plates/plates.go -w 170 -b 10 # if barbel weights 10kgs, by default barbell weight is 20 kgs
go run plates/main.go plates/plates.go -w 170 -c 5 # if use collars that weight 5kg (not each but the full weight of the pair)
# use json
go run plates/main.go plates/plates.go -j ./plates.json
```
