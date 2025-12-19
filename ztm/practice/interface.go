package practice

import "fmt"

// ইন্টারফেস: যে কেউ এই দুটো মেথড করতে পারবে, সে User হিসেবে কাজ করবে
type User1 interface {
	Deposit(amount float64)
	Withdraw(amount float64)  // নাম ঠিক করেছি, বড় W দিয়ে
}

// আরেকটা ইন্টারফেস (পরে ব্যবহার করবো)
type BankManager interface {
	PrintUserDetails()
}

// আমাদের ইউজারের তথ্য রাখার স্ট্রাকচার
type UserInfo struct {
	Name    string
	Balance float64
}

// Deposit মেথড (পয়েন্টার রিসিভার দিয়ে, যাতে Balance চেইঞ্জ হয়)
func (u *UserInfo) Deposit(amount float64) {
	u.Balance += amount
	fmt.Println(amount, "টাকা জমা হয়েছে।")
}

// Withdraw মেথড (নাম বড় W দিয়ে, ইন্টারফেসের সাথে ম্যাচ করতে)
func (u *UserInfo) Withdraw(amount float64) {
	if amount > u.Balance {
		fmt.Println("যথেষ্ট টাকা নেই!")
		return
	}
	u.Balance -= amount
	fmt.Println(amount, "টাকা তোলা হয়েছে।")
}

// Print করার মেথড (BankManager ইন্টারফেসের জন্য)
func (u *UserInfo) PrintUserDetails() {
	fmt.Printf("নাম: %s, ব্যালেন্স: %.2f টাকা\n", u.Name, u.Balance)
}

func Interface() {
	// একটা ইউজার বানাই
	var user User1                     // ইন্টারফেস টাইপের ভেরিয়েবল
	user = &UserInfo{                    // পয়েন্টার দিতে হবে কারণ মেথডগুলো *UserInfo দিয়ে
		Name:    "Rayhan",
		Balance: 10000,
	}

	// ইন্টারফেস দিয়ে কাজ করি
	user.Deposit(500)
	user.Withdraw(200)
	user.Withdraw(20000) // এটা হবে না, টাকা কম

	// এখন Print করতে চাইলে – User ইন্টারফেসে Print নেই!
	// তাই আমরা টাইপ অ্যাসারশন করবো (পরে বুঝাবো)
	if manager, ok := user.(BankManager); ok {
		manager.PrintUserDetails()
	}

	// অথবা সরাসরি
	realUser := user.(*UserInfo) // আমরা জানি এটা UserInfo
	fmt.Printf("চূড়ান্ত ব্যালেন্স: %.2f টাকা\n", realUser.Balance)
}