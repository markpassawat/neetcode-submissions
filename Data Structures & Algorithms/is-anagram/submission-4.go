  func isAnagram(s string, t string) bool {
        if len(s) != len(t) {
                return false
        }

        sb := []byte(s)
        tb := []byte(t)

        sort.Slice(sb, func(i, j int) bool { return sb[i] < sb[j] })
        sort.Slice(tb, func(i, j int) bool { return tb[i] < tb[j] })

        return string(sb) == string(tb)
  }