import signal
import sys
import client

"""KVStore CLI"""


def main():
    print_usage()
    while True:
        line = input("λ ")
        successful = run_command(line)
        if not successful:
            print_usage()


def print_usage():
    print("usage:")
    print("    set key=value")
    print("    get key\n")


def run_command(line) -> bool:
    """Run the command, returning whether the input was well-formed."""
    args = line.split()
    if len(args) != 2:
        return False
    operator = args[0]
    operand = args[1]
    if operator == 'set':
        kv = operand.split(sep='=')
        if len(kv) != 2:
            print(f"Illegal argument to set: '{operand}'")
            return False
        client.set_value(kv[0], kv[1])
        return True
    elif operator == 'get':
        print(client.get_value(operand))
        return True
    else:
        print(f"Illegal command: '{operator}'")
        return False


def exit_gracefully(signal, frame):
    print("\ngood night!")
    client.flush()
    sys.exit(0)


signal.signal(signal.SIGINT, exit_gracefully)

if __name__ == '__main__':
    main()
