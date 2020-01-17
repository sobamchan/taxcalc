package main

import (
    "fmt"
)

// https://financial-field.com/tax/2019/01/25/entry-34511
func main() {
    var income float64

    fmt.Print("Enter Monthly income: ")
    _, err := fmt.Scanf("%f", &income)
    if err != nil {
        fmt.Println(err)
    }
    income *= 12  // To annual income

    // 所得税
    var baseShotokuDeduction, shotokuDeduction float64
    if income < 1800000 {
        baseShotokuDeduction = income * 0.4
        if baseShotokuDeduction < 650000 {
            baseShotokuDeduction = 650000
        }
    } else if 1800000 < income && income < 3600000 {
        baseShotokuDeduction = income * 0.3 + 180000
    } else if 3600000 < income && income < 6600000 {
        baseShotokuDeduction = income * 0.2 + 180000
    } else if 6600000 < income && income < 10000000 {
        baseShotokuDeduction = income * 0.1 + 1200000
    } else {
        baseShotokuDeduction = 2200000
    }
    // 各種所得控除 (生命保険控除, 保険料控除, 勤労学生控除, etc...)
    shotokuDeduction = baseShotokuDeduction + 380000  // 基礎控除

    kazeiShotoku := income - shotokuDeduction

    var shotokuZei float64
    if kazeiShotoku <= 1950000 {
        shotokuZei = kazeiShotoku * 0.05 - 0
    } else if 1950000 < kazeiShotoku && kazeiShotoku <= 3300000 {
        shotokuZei = kazeiShotoku * 0.1 - 97500
    } else if 3300000 < kazeiShotoku && kazeiShotoku <= 6950000 {
        shotokuZei = kazeiShotoku * 0.2 - 427500
    } else if 6950000 < kazeiShotoku && kazeiShotoku <= 9000000 {
        shotokuZei = kazeiShotoku * 0.23 - 636000
    } else if 9000000 < kazeiShotoku && kazeiShotoku <= 18000000 {
        shotokuZei = kazeiShotoku * 0.33 - 1536000
    } else if 18000000 < kazeiShotoku && kazeiShotoku <= 40000000 {
        shotokuZei = kazeiShotoku * 0.4 - 2796000
    } else {
        shotokuZei = kazeiShotoku * 0.45 - 4796000
    }

    // 住民税
    zyuminDeduction := baseShotokuDeduction + 330000
    zyuminKazeiShotoku := income - zyuminDeduction
    zyuminZei := zyuminKazeiShotoku * 0.1

    fmt.Printf("年収:     %f\n", income)
    fmt.Printf("所得税: %f\n", shotokuZei)
    fmt.Printf("住民税: %f\n", zyuminZei)
    fmt.Printf("所得税 + 住民税: %f\n", shotokuZei + zyuminZei)
    fmt.Printf("年収 - (所得税 + 住民税): %f\n", income - (shotokuZei + zyuminZei))

}
