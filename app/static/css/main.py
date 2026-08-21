p: str = './technologies'


data: list[str] = []

with open(p + '.css', 'r') as doc:
    for line in doc:
        if '{' in line:
            line = line.split('{')[0].strip()
            data.append(line)



if ( slen := len(set(data)) )!= ( lslen := len(data) ):
    print('items duplicateds: ', slen, lslen)


def _() -> None:
    to_mod: list[str] = []

    with open(p + '-dark.css', 'r') as doc:
        to_mod = doc.readlines()
        for line in doc:
            if '{' in line:
                item = line.split('{')[0].strip()

                if not item in data:
                    print('new item not exists in original: ', item)

        else: print('all classes exists')


    with open(p + '-dark.css', 'w') as doc:
        doc.writelines(
            line
            for line in to_mod
        )
