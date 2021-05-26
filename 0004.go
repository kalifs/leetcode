func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
    len1 := len(nums1)
    len2 := len(nums2)
    done1 := false
    done2 := false

    idx :=  (len1 + len2) / 2
    c1 := 0
    c2 := 0
    var n1, n2 int
    results := make([]int, 2)

    for i:= 0; i<= idx; i++ {
        if c1 < len1 {
            n1 = nums1[c1]
        } else {
            done1 = true
        }
        if c2 < len2 {
            n2 = nums2[c2]
        } else {
            done2 = true
        }


        if (!done1 && n1 < n2) || (done2 && !done1) {
            results[i%2] = nums1[c1]
            c1 += 1
        }
        if (!done2 && n2 <= n1) || (done1 && !done2) {
            results[i%2] = nums2[c2]
            c2 += 1
        }
    }

    if len1 + len2 > 1 {
        if (len1 + len2) % 2 == 0 {
            return (float64(results[0]) + float64(results[1])) / 2.0
        } else {
            return float64(results[((len1 + len2) / 2) % 2])
        }
    } else {
        return float64(results[0])
    }
}
