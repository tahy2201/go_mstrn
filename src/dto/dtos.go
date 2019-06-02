package dto

import (
	"time"
)

type AddTrainData struct {
	ExerciseID     int     `json:"exercise_id"`
	UserID         string  `json:"user_id"`
	SetCount       int     `json:"set_count"`
	ExerciseWeight float32 `json:"exercise_weight"`
	ExerciseTimes  int     `json:"exercise_times"`
	ExerciseMemo   string  `json:"exercise_memo"`
}

// TUserProfile ユーザプロフィール
type TUserProfile struct {
	UserID      string    `json:"user_id"`
	UserName    string    `json:"user_name"`
	Birthday    time.Time `json:"birthday"`
	MailAddress string    `json:"mail_address"`
	Password    string    `json:"password"`
	Profile     string    `json:"profile"`
	RegistData  time.Time `json:"regist_data"`
	CreateUser  string    `json:"create_user"`
	CreateDate  time.Time `json:"create_date"`
	UpdateUser  string    `json:"update_user"`
	UpdateDate  time.Time `json:"update_date"`
}

func (m *TUserProfile) TableName() string {
	return "t_user_profile"
}

// TTrainRecord トレーニング記録
type TTrainRecord struct {
	TrainID        int       `json:"train_id"`
	ExerciseID     int       `json:"exercise_id"`
	UserID         string    `json:"user_id"`
	SetCount       int       `json:"set_count"`
	ExerciseWeight float32   `json:"exercise_weight"`
	ExerciseTimes  int       `json:"exercise_times"`
	ExerciseMemo   string    `json:"exercise_memo"`
	CreateUser     string    `json:"create_user"`
	CreateDate     time.Time `json:"create_date"`
	UpdateUser     string    `json:"update_user"`
	UpdateDate     time.Time `json:"update_date"`
}

func (m *TTrainRecord) TableName() string {
	return "t_train_record"
}

// TExerciseMaster 種目マスタ
type TExerciseMaster struct {
	ExerciseID   int       `json:"exercise_id"`
	ExerciseName string    `json:"exercise_name"`
	Description  string    `json:"description"`
	MachineMaker string    `json:"machine_maker"`
	CreateUser   string    `json:"create_user"`
	CreateDate   time.Time `json:"create_date"`
	UpdateUser   string    `json:"update_user"`
	UpdateDate   time.Time `json:"update_date"`
}

func (m *TExerciseMaster) TableName() string {
	return "t_exercise_master"
}

// TUserPhysicalData 身体データ
type TUserPhysicalData struct {
	UserID            string    `json:"user_id"`
	DataNum           string    `json:"data_num"`
	Height            float32   `json:"height"`
	Weight            float32   `json:"weight"`
	BodyFatPercentage float32   `json:"body_fat_percentage"`
	BloodPressureMin  float32   `json:"blood_pressure_min"`
	BloodPressureMax  float32   `json:"blood_pressure_max"`
	Pulse             float32   `json:"pulse"`
	CreateUser        string    `json:"create_user"`
	CreateDate        time.Time `json:"create_date"`
	UpdateUser        string    `json:"update_user"`
	UpdateDate        time.Time `json:"update_date"`
}

func (m *TUserPhysicalData) TableName() string {
	return "t_user_physical_data"
}
