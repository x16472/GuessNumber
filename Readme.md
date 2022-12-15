# 猜數字0-100
## C++轉換Golang

-   原程式碼
```c++
#include <iostream>
using namespace std;

int main()
{
    int answer = 27, guess = -1;
    while (guess!=answer) {
        cout << "請輸入1-100間的數字:";
        cin >> guess;
        if (guess==answer) {
            cout << "恭喜你猜對了";
            break;
        }
        else if(guess<answer)
        {
            cout << "太小了" << endl;
            if (answer-guess <= 5) {
                cout << "不過很接近了" << endl;
            }
        }
        else if (answer > guess) {
            cout << "太大了" << endl;
            if (answer-guess >= 5) {
                cout << "不過很接近了" << endl;
            }
        }
    }
    return 0;
}
```# GuessNumber
# GuessNumber
