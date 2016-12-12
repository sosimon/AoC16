import re

def supports_tls(ip):
    abba_pattern = re.compile(r'(\w)(\w)\2\1')
    hypernet_pattern = re.compile(r'\[(.*?)\]')
    results = abba_pattern.findall(ip)
    if results:
        if results[0][0] == results[0][1]:
            return False
        hypernet_seqs = hypernet_pattern.findall(ip)
        for seq in hypernet_seqs:
            if abba_pattern.findall(seq):
                return False
        return True
    else:
        return False

def supports_ssl(ip):
    parse_pattern = re.compile(r'(.*?)(\[(.*?)\](\w*)?)+?')
    aba_pattern = re.compile(r'(?=(.)(.)\1)')
    matches = parse_pattern.findall(ip)
    supernet = []
    hypernet = []
    if matches:
        for m in matches:
            if m[0]:
                supernet.append(m[0])
            supernet.append(m[-1])
            hypernet.append(m[-2])
    for s_seq in supernet:
        results = aba_pattern.findall(s_seq)
        if results:
            for r in results:
                if r[0] == r[1]:
                    continue
                bab_str = ''.join(r[1] + r[0] + r[1])
                for h_seq in hypernet:
                    if bab_str in h_seq:
                        return True
    return False

if __name__ == "__main__":
    count = 0
    with open("input", 'r') as f:
        for line in f:
            ip = line.rstrip("\n")
            print ip
            if supports_ssl(ip):
                print "True"
                count += 1
    print "Number of IPs supporting SSL: {0}".format(count)
