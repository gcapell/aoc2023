digits = [
    'zero', 'one', 'two', 'three', 'four', 'five', 'six',
    'seven', 'eight', 'nine',
    '0', '1', '2', '3', '4', '5', '6', '7', '8', '9'
]

def score(line):
    earliest = ('',-1)
    latest = ('', -1)
    for d in digits:
        n = line.find(d)
        if n == -1:
            continue
        if earliest[0] == '' or n < earliest[1]:
            earliest = (d,n)
        n = line.rfind(d)
        if n == -1:
            continue
        if latest[0] == '' or n > latest[1]:
            latest = (d,n)
    return to_n(earliest[0]) * 10 + to_n(latest[0])
    
def to_n(s):
    if (n := digits.index(s)) < 10:
        return n
    return int(s)
        
    
def main():
    lines = open('input.txt').readlines()
    print (sum(map(score,lines)))

    
if __name__ == '__main__':
    main()