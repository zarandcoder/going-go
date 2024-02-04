name="Vadim"

def introduce(ur_name):
    a = 5
    b = 10
    print(f'a={a}; b={b}.....and now swap...')
    a,b = b,a
    print(f'a={a}; b={b}.....finish')
    print(f"Hello {ur_name}")
    print('How are you doing today?')

introduce(name)
