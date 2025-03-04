#include <bits/stdc++.h>
using namespace std;

int main()
{
    int arr[14] = {2, 4, 6, 4, 2, 3, 4, 1, 5, 6, 6, 6, 6, 6};
    int A = 2, B = 4, C = 6;
    // Time limit = 10
    map<int, int> mp;
    for (int i = 0; i < 14; i++)
    {
        int people = arr[i];
        if (people == 1)
            people = 2;
        if (people == 3)
            people = 4;
        if (people == 5)
            people = 6;
        mp[people]++;
    }
    int mn = 1e6 + 10;
    for (auto i : mp)
    {
        mn = min(mn, i.second);
    }
    mp[A] -= mn;
    mp[B] -= mn;
    mp[C] -= mn;

    int unsetted = 0;
    while (1)
    {

        if (mp[A] == 0 && mp[B] == 0 && mp[C] == 0)
            break;
        if (mp[C] > 0)
        {
            mp[C]--;
        }
        else if (mp[B] > 0)
        {
            mp[B]--;
        }
        else if (mp[A] > 0)
        {
            mp[A]--;
        }
        else
        {
            unsetted++;
        }
        if (mp[B] > 0)
        {
            mp[B]--;
        }
        else if (mp[A] > 0)
        {
            mp[A]--;
        }
        else
        {
            unsetted++;
        }

        if (mp[A] > 0)
        {
            mp[A]--;
        }
        else
        {
            unsetted++;
        }
    }

    cout << "Unsetted Table :" << unsetted << endl;
    return 0;
}


// Time Complexity of this code is O(n)