name="Vadim"

def demo_swap(a=1, b=2):
    print('Demo swap')
    print(f'a={a}; b={b}....aand now magic is about to happen!')
    a,b = b,a
    print(f'Finish with magic...results are a={a}, b={b}')
    

def introduce(ur_name):
    print(f"Hello {ur_name}")
    print('How are you doing today?')

demo_swap()
introduce(name)
