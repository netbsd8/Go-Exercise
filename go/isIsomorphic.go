package main

func isIsomorphic(s string, t string) bool {
	length := len(s)
	m_st := make(map[byte]byte)
	m_ts := make(map[byte]byte)

	for i:=0; i<length; i++ {
		/*
		if s[i] == t[i] {
			return false
		}
		*/

		ans_st, ok_st := m_st[s[i]]
		ans_ts, ok_ts := m_ts[t[i]]
		if !ok_st && !ok_ts {
			m_st[s[i]] = t[i]
			m_ts[t[i]] = s[i]
		} else {
			if ok_st && ok_ts && ans_st == t[i] && ans_ts == s[i] {
				continue
			}
			return false	
		} 
	}

	return true
}