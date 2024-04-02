/**
An implementation of Weighted Union + Collapsiong Find for n nodes.
*/

type DisjointSet struct{
    Parent []int
}

func (ds *DisjointSet) Init(n int){
    arr:=make([]int, n)
    for i:=range arr{
        arr[i] = -1
    }
    ds.Parent = arr
}

func (ds *DisjointSet) Find(a int) int{
    if ds.Parent[a] < 0{
        return a
    }
    ds.Parent[a] = ds.Find(ds.Parent[a])
    return ds.Parent[a]
}

func (ds *DisjointSet) Union(a, b int) bool{
    pa,pb:=ds.Find(a),ds.Find(b)
    if pa == pb{return false}
    if ds.Parent[pa]<ds.Parent[pb]{
        ds.Parent[pa]+=ds.Parent[pb]
        ds.Parent[pb] = pa
    }else{
        ds.Parent[pb]+=ds.Parent[pa]
        ds.Parent[pa]=pb
    }
    return true
}
