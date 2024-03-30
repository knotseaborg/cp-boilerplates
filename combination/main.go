/**
  nCr with modulo 10e9+7. Warning time complexity is O(n^2).
  Pay attention to edge cases.

  Can probably be optimized to O(n*r), will do later.
*/

const P = 1000_000_007

func combination(n, r int) int{
    dp:=make([]int, n+1)
    dp[0]=1
    for i:=1;i<=n;i++{
        for j:=i;j>=1;j--{
            dp[j]+=dp[j-1]
            dp[j]%=P
        }
    }
    return dp[r-1]
}
