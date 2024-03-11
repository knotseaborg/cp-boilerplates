/**
  This is a template for rolling hash, which can be used to efficiently compare substrings in a large string. 

  Given:
    s = "This_is_a_dummy_string"
             i<---->j<---->k 
    if (hash[i]*Pow[j])%P == (P+hash[k]-hash[j-1])%P {substrings are the same}
  
  **Conflicts are possible.

  P must be a large number, which limits number of unique hashes possible. Larger numbers reduce conflicts.
  M can ideally be a power of 2. Must be greater than the number unique characters in the strings. 
*/
const M, P = 32, 1000_000_007

// Stores powers of M
var Pow []int

// Generates powers of M and stores in Pow
func InitPow(n int){
    Pow=make([]int, n+1)
    Pow[0]=1
    for i:=1;i<=n;i++{
        Pow[i] = M*ans[i-1]
        Pow[i]%=P
    }
}

// Generates Rolling Hash of string at every index
func HashGen(s string) []int{
    n:=len(s)
    ans:=make([]int, n)
    ans[0]=int(s[0]-'a')+1
    for i:=1;i<n;i++{
        ans[i] = ans[i-1]+(int(s[i]-'a'+1)*Pow[i])%P
        ans[i]%=P
    }
    return ans
}
