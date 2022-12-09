def generateAlphabetList():
  """
  This Function Generate Alphabet List which will be used on Key Table Generation
  """
  alphabetList = []
  for letter in range(ord('a'), ord('z')+1):
    if(chr(letter) != 'j'):
      alphabetList.append(chr(letter))
  return alphabetList

def generateKeyTable(key):
  """
  This Function Generate Key Table which takes an input of key
  """
  keyList = []
  keyTable = []
  alphabetList = generateAlphabetList()
  # Append Key to Table
  for letter in key:
    if letter not in keyList and letter != 'j':
      keyList.append(letter)
  # Append Alphabet to Table
  for letter in alphabetList:
    if letter not in keyList and letter != 'j':
      keyList.append(letter)
  # Turn into 2D 5x5 list
  while keyList != []:
    keyRow = []
    for i in range(0,5):
      if(keyList[0] is not None):
        keyRow.append(keyList.pop(0))
    keyTable.append(keyRow)
  return keyTable

def search(keyTable, char):
  """
  This Function searches the index of the character
  """
  for i, x in enumerate(keyTable):
    for a, b in enumerate(x):
      if char == b:
        return i,a

def next_column(keyTable, char):
  """
  This Functions Does the next column algorithm if its on same column
  """
  x, y = search(keyTable, char)
  yNext = (y+1) % 5
  return keyTable[x][yNext]

def next_row(keyTable, char):
  """
  This Functions Does the next row algorithm if its on same row
  """
  x, y = search(keyTable, char)
  xNext = (x+1) % 5
  return keyTable[xNext][y]

def encrypt(keyTable, segment):
  first = segment[0]
  second = segment[1]
  x_first, y_first = search(keyTable, first)
  x_second, y_second = search(keyTable, second)
  if x_first == x_second:
    return "{}{}".format(
        next_column(keyTable, first),
        next_column(keyTable, second),
    )
  elif y_first == y_second:
    return "{}{}".format(
        next_row(keyTable, first),
        next_row(keyTable, second),
    )
  else:
    # This Return The swapping algorithm if its on different row and column
    return "{}{}".format(keyTable[x_first][y_second], keyTable[x_second][y_first])

def get_bool(prompt):
    """
    This Functions is used to user input true or false
    """
    while True:
      try:
          return {"true":True,"false":False,1:True,0:False,"yes":True,"no":False}[input(prompt).lower()]
      except KeyError:
          print("Invalid input please enter True or False!")

def main():
  separator = '\n-------------------------------------'
  loop = 0
  print("PlayFair Cipher Encryption")

  # Key and PlainText Input
  while True:
    print(separator)
    if loop:
      key = input('Enter Key Again: ').replace(" ", "").lower()
      plaintext = input('Enter PlainText Again: ').replace(" ", "").lower()
    else:
      key = input('Enter Key: ').replace(" ", "").lower()
      plaintext = input('Enter PlainText: ').replace(" ", "").lower()

    # Helper Code to Visualize Process
    print("Key:", key)
    keyTable = generateKeyTable(key)
    print("Key Table")
    for k in keyTable:
      print(k)
    print()
    print("Plaintext:", plaintext)

    # Running Encryption
    segments =  []
    ignore = False
    for i in range(len(plaintext)):
      if ignore:
        ignore = False
        continue
      nextIndex = i+1
      if nextIndex == len(plaintext):
        plaintext += "z"
      c = plaintext[i]
      cn = plaintext[nextIndex]
      if c != cn:
        segments.append(encrypt(keyTable, "{}{}".format(c, cn)))
        ignore = True
      else:
        segments.append(encrypt(keyTable, "{}{}".format(c, "x")))

    # Join The Segments into one CipherText
    cipher = "".join(segments)
    print("Cipher:", cipher)

    # Retry the PlayFair Cipher
    isAgain = get_bool("Would You Like to Try Again? [True or False]:")
    if not isAgain:
      print(separator)
      print("Thanks For Trying out our code")
      break
    loop = loop + 1

if __name__ == "__main__":
    main()