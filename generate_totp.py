import hmac
import hashlib
import time
import struct

def generate_totp(secret, interval=30):
    key = secret.encode()
    msg = struct.pack('>Q', int(time.time()) // interval)
    hmac_hash = hmac.new(key, msg, hashlib.sha512).digest()
    offset = hmac_hash[-1] & 0x0F
    code = (struct.unpack('>I', hmac_hash[offset:offset + 4])[0] & 0x7FFFFFFF) % 10000000000
    return str(code).zfill(10)

email = "kirill.sorokin@icloud.com"
shared_secret = email + "HENNGECHALLENGE003"
totp_password = generate_totp(shared_secret)
print(totp_password)