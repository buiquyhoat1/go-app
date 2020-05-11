package seed

import (
	"log"

	"github.com/jinzhu/gorm"
	"go_api/models"
)

var users = []models.User{
	models.User{
		Name: "Wayner Rooney",
	},
	models.User{
		Name: "David Beckham",
	},
	models.User{
		Name: "Francesco Totti",
	},
	models.User{
		Name: "Paul Scholes",
	},
	models.User{
		Name: "Ryan Giggs",
	},
}

var posts = []models.Post{
	models.Post{
		Title: "一期一会",
		Content: "「一期一会」とは生に一度だけの機会。生涯に一度限りであること。生涯に一回しかないと考えて、そのことに専念する意。もと茶道の心得を表した語で、どの茶会でも一生に一度のものと心得て、主客ともに誠意を尽くすべきことをいう。▽千利休の弟子宗二の『山上宗二記やまのうえそうじき』に「一期に一度の会」とあるのによる。「一期」は仏教語で、人が生まれてから死ぬまでの間の意。",
	},
	models.Post{
		Title: "雨降って地固まる",
		Content: "「雨降って地固まる」の意味は「いざこざやアクシデントの後は、かえって前よりも良い状態になること」です。悪い出来事やゴタゴタが起こった後は、前より基礎が固まって事態が良くなること・結束力が強くなることを表します。",
	},
	models.Post{
		Title: "水に流す",
		Content: "「水に流す」とは、お互いの間にあったいざこざや気まずさを、なかったことにするという意味のことわざです。 つまり、過去のトラブルなどで悪くなった仲を一度リセットして、再度新しく関係性をはじめることを指します。 条件付きで相手を許すときに「水に流す」と言うこともあります。 ... 同じ意味で、「水にする」「水になす」とも言います。",
	},
	models.Post{
		Title: "猫に小判",
		Content: "ことわざ （猫に小判を与えても、その価値を知らない猫にとっては何の意味もないことから）どんな立派なものでも、価値がわからない者にとっては、何の値打ちもないものであるというたとえ。",
	},
	models.Post{
		Title: "七転八起",
		Content: "「七転八起」の意味は「何回失敗しても、諦めずに立ち上がって奮闘すること」です。いくら落ち込んでいたとしても、屈せずに再び挑戦することで良い事が起きることを表します。また、「人生は波乱万丈である」という意味も含まれます。",
	},
}

var blogs = []models.Blog{
	models.Blog{
		Title: "一期一会",
		Content: "「一期一会」とは生に一度だけの機会。生涯に一度限りであること。生涯に一回しかないと考えて、そのことに専念する意。もと茶道の心得を表した語で、どの茶会でも一生に一度のものと心得て、主客ともに誠意を尽くすべきことをいう。▽千利休の弟子宗二の『山上宗二記やまのうえそうじき』に「一期に一度の会」とあるのによる。「一期」は仏教語で、人が生まれてから死ぬまでの間の意。",
		Author: "Nguyen Van A",
		Category: "A",
		Status: 0,
	},
	models.Blog{
		Title: "雨降って地固まる",
		Content: "「雨降って地固まる」の意味は「いざこざやアクシデントの後は、かえって前よりも良い状態になること」です。悪い出来事やゴタゴタが起こった後は、前より基礎が固まって事態が良くなること・結束力が強くなることを表します。",
		Author: "Nguyen Van B",
		Category: "B",
		Status: 0,
	},
	models.Blog{
		Title: "水に流す",
		Content: "「水に流す」とは、お互いの間にあったいざこざや気まずさを、なかったことにするという意味のことわざです。 つまり、過去のトラブルなどで悪くなった仲を一度リセットして、再度新しく関係性をはじめることを指します。 条件付きで相手を許すときに「水に流す」と言うこともあります。 ... 同じ意味で、「水にする」「水になす」とも言います。",
		Author: "Nguyen Van C",
		Category: "C",
		Status: 0,
	},
	models.Blog{
		Title: "猫に小判",
		Content: "ことわざ （猫に小判を与えても、その価値を知らない猫にとっては何の意味もないことから）どんな立派なものでも、価値がわからない者にとっては、何の値打ちもないものであるというたとえ。",
		Author: "Nguyen Van D",
		Category: "D",
		Status: 0,
	},
	models.Blog{
		Title: "七転八起",
		Content: "「七転八起」の意味は「何回失敗しても、諦めずに立ち上がって奮闘すること」です。いくら落ち込んでいたとしても、屈せずに再び挑戦することで良い事が起きることを表します。また、「人生は波乱万丈である」という意味も含まれます。",
		Author: "Nguyen Van E",
		Category: "A",
		Status: 1,
	},
}

func Load(db *gorm.DB) {
	// check_users := []*models.User{}
	// var count int
	// db.Table("users").Count(&count)
	if(db.Debug().First(&models.User{}).RecordNotFound()) {
		err := db.Debug().DropTableIfExists(&models.Post{}, &models.User{}, &models.Blog{}).Error
		if err != nil {
			log.Fatalf("cannot drop table: %v", err)
		}
		err = db.Debug().AutoMigrate(&models.User{}, &models.Post{}, &models.Blog{}).Error
		if err != nil {
			log.Fatalf("cannot migrate table: %v", err)
		}

		/*
			err = db.Debug().Model(&models.Post{}).AddForeignKey("author_id", "users(id)", "cascade", "cascade").Error
			if err != nil {
				log.Fatalf("attaching foreign key error: %v", err)
			}
		*/

		for i, _ := range users {
			err = db.Debug().Model(&models.User{}).Create(&users[i]).Error
			if err != nil {
				log.Fatalf("cannot seed users table: %v", err)
			}
			posts[i].UserID = users[0].ID

			err = db.Debug().Model(&models.Post{}).Create(&posts[i]).Error
			if err != nil {
				log.Fatalf("cannot seed posts table: %v", err)
			}
		}

		for i, _ := range blogs {
			err = db.Debug().Model(&models.Blog{}).Create(&blogs[i]).Error
			if err != nil {
				log.Fatalf("cannot seed blogs table: %v", err)
			}
		}
	}
}
