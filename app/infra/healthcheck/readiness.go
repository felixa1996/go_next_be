package healthcheck

var mongoReadiness bool = false
var awsSessionReadiness bool = false

func SetMongoReadiness(status bool) {
	mongoReadiness = status
}

func GetMongoReadiness() bool {
	return mongoReadiness
}

func SetAwsSessionReadiness(status bool) {
	awsSessionReadiness = status
}

func GetAwsSessionReadiness() bool {
	return awsSessionReadiness
}
