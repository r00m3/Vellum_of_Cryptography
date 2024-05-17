
┏━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┓


        Vellum_of_Cryptography v0.0.1
                        by r00m3

        ・────────── ⧝ ──────────・

        ▸ Uncomplicated way to get you familiar with concepts of cryptography.
        ▸ To use these concepts in real world, look for other tools by r00m3.


□ ────────────────────────────────────────────────────────────────────────────────────────────────── □


        [ Confidentiality ] - hide the message.
        [ Integrity ]       - message is unchanged.
        [ Authentication ]  - message belongs to actual sender.


□ ────────────────────────────────────────────────────────────────────────────────────────────────── □


        ▸ Vellum_of_Cryptography :
                ✓ intentionally uses very abstract examples.
                ✓ intentionally skips complex implementation details.
                ✓ is not complete.
                ✓ is in early version.
                ✓ might and probably do contain inaccurate information.
                ✓ will be improved and fixed in further releases.


□ ────────────────────────────────────────────────────────────────────────────────────────────────── □


        ▸ Notation used in Vellum_of_Cryptography :

                [ Function description ]
        FUNCTION_NAME(FUNCTION_PARAMETERS) -> FUNCTION_RETURN


□ ────────────────────────────────────────────────────────────────────────────────────────────────── □


        ▸ Some terms in Vellum_of_Cryptography are used interchangeably :

        SECRET_MESSAGE ⇄ ENCRYPTED_MESSAGE ⇄ ENCRYPTED_DOCUMENT ⇄ ENCRYPTED_FILE


┗━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┛


┏━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┓


        Hash functions

        ・────────── ⧝ ──────────・


        ▸ Create file.txt
        ▸ Open file.txt
        ▸ Type letter a
        ▸ Save and close file.txt
        ▸ Open terminal in same folder as file.txt

        ▸ On Linux type :
                sha256sum file.txt
        ▸ On MacOS type :
                shasum file.txt -a 256
        ▸ On Windows type :
                get-filehash file.txt -algo sha256

        ▸ You will get :
                ca978112ca1bbdcafac231b39a23dc4da786eff8147c4e72b9807785afee48bb
        ▸ You just calculated sha256 HASH or checksum of file.txt


        ・────────── ⧝ ──────────・


        ▸ Change content in file.txt by replacing letter a with letter b
        ▸ Calculate file.txt HASH

        ▸ You will get :
                3e23e8160039594a33894f6564e1b1348bbd7a0088d42c4acb73eeaed59c009d


        ・────────── ⧝ ──────────・


        ▸ Every file has it's unique HASH.
        ▸ If you change even a bit of that file, HASH will also change.
        ▸ You know that file.txt with letter a inside it will produce this HASH :
                ca978112ca1bbdcafac231b39a23dc4da786eff8147c4e72b9807785afee48bb
        ▸ You know this because you just calculated it.
        ▸ Everyone can do the same. Create file.txt and calculate the HASH.
        ▸ What you CAN'T do :
                ✓ know what file with what content produce
                ca978112ca1bbdcafac231b39a23dc4da786eff8147c4e72b9807785afee48bb
        ▸ You can NOT guess file knowing it's HASH.
        ▸ You CAN calculate HASH if you have the file.


□ ────────────────────────────────────────────────────────────────────────────────────────────────── □


        PERSON_A creates FILE_A and calculates HASH_A_1.
        PERSON_A puts FILE_A on the internet.
        PERSON_A puts HASH_A_1 on his PRIVATE_WEBSITE.
        PERSON_B knows PRIVATE_WEBSITE belongs to PERSON_A.
        PERSON_C changes content of FILE_A and says it's same FILE_A from PERSON_A.
        PERSON_B downloads FILE_A and calculates HASH_A_2.
        PERSON_B compares HASH_A_2 which they just calculated
        with HASH_A_1 from PRIVATE_WEBSITE.
        PERSON_B sees that HASH_A_1 and HASH_A_2 does not match.
        PERSON_B now knows that even if FILE_A looks the same,
        someone on the internet changed it.
        PERSON_B now knows that it is not safe to use FILE_A.


┗━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┛



┏━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┓


        Authenticated encryption [ AE ]

        ・────────── ⧝ ──────────・

        ▸ PERSON_A generates PERSON_A_SECRET_KEY.
        ▸ PERSON_A encrypts PERSON_A_SECRET_MESSAGE with PERSON_A_SECRET_KEY.
        ▸ Only ones which have PERSON_A_SECRET_KEY can :
                ✓ encrypt PERSON_A_SECRET_MESSAGE.
                ✓ decrypt PERSON_A_SECRET_MESSAGE.
                ✓ modify PERSON_A_SECRET_MESSAGE.


□ ────────────────────────────────────────────────────────────────────────────────────────────────── □


                [ Generate PERSON_A_SECRET_KEY ]
        GENERATE_PERSON_A_SECRET_KEY() -> PERSON_A_SECRET_KEY

                [ Encrypt PERSON_A_SECRET_MESSAGE ]
        ENCRYPT(PERSON_A_SECRET_KEY, NONCE, UNENCRYPTED_MESSAGE) -> PERSON_A_SECRET_MESSAGE

                [ NONCE = a number that can be used only once. ]
        NONCE is not required to be secret.
        You must not use same NONCE twice when encrypting with same SECRET_KEY.

                [ Decrypt PERSON_A_SECRET_MESSAGE ]
        DECRYPT(PERSON_A_SECRET_KEY, NONCE, PERSON_A_SECRET_MESSAGE) -> UNENCRYPTED_MESSAGE


□ ────────────────────────────────────────────────────────────────────────────────────────────────── □


        ▸ Authenticated encryption with associated data [ AEAD ]
        ▸ Same as [ AE ], but some data(associated data) can be left in plaintext(unencrypted).

                [ Generate PERSON_A_SECRET_KEY ]
        GENERATE_PERSON_A_SECRET_KEY() -> PERSON_A_SECRET_KEY

                [ Encrypt PERSON_A_SECRET_MESSAGE with associated data ]
        ENCRYPT(PERSON_A_SECRET_KEY, NONCE, UNENCRYPTED_MESSAGE, UNENCRYPTED_ADDITIONAL_DATA) ->
        PERSON_A_SECRET_MESSAGE, UNENCRYPTED_ADDITIONAL_DATA

                [ Decrypt PERSON_A_SECRET_MESSAGE with associated data ]
        DECRYPT(PERSON_A_SECRET_KEY, NONCE, PERSON_A_SECRET_MESSAGE, UNENCRYPTED_ADDITIONAL_DATA) ->
        UNENCRYPTED_MESSAGE, UNENCRYPTED_ADDITIONAL_DATA

        ▸ Note that if someone without PERSON_A_SECRET_KEY tries to change UNENCRYPTED_ADDITIONAL_DATA
        DECRYPT() function will not proceed and warn user that data was changed.


┗━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┛



┏━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┓


        Symmetric encryption

        ・────────── ⧝ ──────────・

        ▸ Generate SECRET_KEY.
        ▸ Encrypt the SECRET_MESSAGE with SECRET_KEY.
        ▸ SECRET_MESSAGE can be decrypted only with SECRET_KEY.


□ ────────────────────────────────────────────────────────────────────────────────────────────────── □


        PERSON_A wants to send a SECRET_MESSAGE to PERSON_B.
        PERSON_A wants to be sure :
                ✓ only PERSON_B can read the SECRET_MESSAGE.
        PERSON_A generates a SECRET_KEY.
        PERSON_A uses SECRET_KEY to encrypt a SECRET_MESSAGE.
        PERSON_A shares a SECRET_KEY and SECRET_MESSAGE with PERSON_B.
        PERSON_B uses SECRET_KEY to read a SECRET_MESSAGE.


□ ────────────────────────────────────────────────────────────────────────────────────────────────── □


        PERSON_A wants to receive a SECRET_MESSAGE from PERSON_B.
        PERSON_A generates SECRET_KEY.
        PERSON_A sends SECRET_KEY to PERSON_B.
        PERSON_B uses SECRET_KEY to encrypt a SECRET_MESSAGE.
        PERSON_B sends SECRET_MESSAGE to PERSON_A.
        PERSON_A uses SECRET_KEY to read the SECRET_MESSAGE.


□ ────────────────────────────────────────────────────────────────────────────────────────────────── □


                [ Generate SECRET_KEY ]
        GENERATE_SECRET_KEY() -> SECRET_KEY

                [ Encrypt message with SECRET_KEY ]
        ENCRYPT(SECRET_KEY, UNENCRYPTED_MESSAGE) -> ENCRYPTED_MESSAGE

                [ Decrypt ]
        DECRYPT(SECRET_KEY, ENCRYPTED_MESSAGE) -> UNENCRYPTED_MESSAGE


┗━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┛


┏━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┓


        Asymmetric encryption [ Public-key cryptography ]

        ・────────── ⧝ ──────────・

        ▸ Generate SECRET_PRIVATE_KEY. Don't share it with anyone.
        ▸ Compute NON_SECRET_PUBLIC_KEY from SECRET_PRIVATE_KEY. You can share it with everyone.
        ▸ If you sign a NON_SECRET_MESSAGE with your SECRET_PRIVATE_KEY,
          Everyone who has your NON_SECRET_PUBLIC_KEY can verify that :
                ✓ you are the one who wrote NON_SECRET_MESSAGE.
                ✓ no body changed NON_SECRET_MESSAGE.
        ▸ Everyone who has your NON_SECRET_PUBLIC_KEY can write you a SECRET_PRIVATE_MESSAGE.
        ▸ Only you, who has SECRET_PRIVATE_KEY will be able to read a SECRET_PRIVATE_MESSAGE.


□ ────────────────────────────────────────────────────────────────────────────────────────────────── □


        PERSON_A wants to put the NON_SECRET_MESSAGE to the internet.
        PERSON_A wants to make sure :
                ✓ everyone on the internet knows PERSON_A and not someone else wrote the NON_SECRET_MESSAGE.
                ✓ everyone on the internet knows that NON_SECRET_MESSAGE is unmodified.
                ✓ it is exact NON_SECRET_MESSAGE which PERSON_A wrote.

        PERSON_A creates a PERSON_A_PRIVATE_KEY.
        PERSON_A computes a PERSON_A_PUBLIC_KEY from PERSON_A_PRIVATE_KEY.
        PERSON_A signs NON_SECRET_MESSAGE with PERSON_A_PRIVATE_KEY.
        PERSON_A now has NON_SECRET_MESSAGE_SIGNATURE.
        PERSON_A puts PERSON_A_PUBLIC_KEY on the internet.
        PERSON_A puts NON_SECRET_MESSAGE_SIGNATURE on the internet.
        PERSON_A puts NON_SECRET_MESSAGE on the internet.
        Everyone on the internet can download PERSON_A_PUBLIC_KEY.
        Everyone on the internet can download NON_SECRET_MESSAGE_SIGNATURE.
        Everyone on the internet can download NON_SECRET_MESSAGE.

        Everyone on the internet can use a PROGRAM where they put :
                ✓ PERSON_A_PUBLIC_KEY,
                ✓ NON_SECRET_MESSAGE_SIGNATURE,
                ✓ NON_SECRET_MESSAGE.
        And program tells if NON_SECRET_MESSAGE was signed with PERSON_A_PRIVATE_KEY.
        If PROGRAM confirm, that NON_SECRET_MESSAGE was signed with PERSON_A_PRIVATE_KEY :
                ✓ everyone on the internet knows PERSON_A and not someone else wrote NON_SECRET_MESSAGE.
                ✓ everyone on the internet knows that NON_SECRET_MESSAGE is unmodified.
                It is exact NON_SECRET_MESSAGE which PERSON_A wrote.
                        ✓ because only PERSON_A knows PERSON_A_PRIVATE_KEY.
                        ✓ and only PERSON_A_PRIVATE_KEY can be used to sign NON_SECRET_MESSAGE.
                        ✓ PERSON_A_PUBLIC_KEY can be used to verify message, but not to sign it.


□ ────────────────────────────────────────────────────────────────────────────────────────────────── □


        PERSON_A wants to receive a PERSON_A_SECRET_MESSAGE dedicated only to PERSON_A.
        PERSON_A wants to be sure :
                ✓ PERSON_A_SECRET_MESSAGE is dedicated only for PERSON_A and no one else can read it.
                ✓ anyone can write a PERSON_A_SECRET_MESSAGE dedicated only for PERSON_A.

        PERSON_A creates a PERSON_A_PRIVATE_KEY.
        PERSON_A computes a PERSON_A_PUBLIC_KEY from PERSON_A_PRIVATE_KEY.
        PERSON_A puts PERSON_A_PUBLIC_KEY on the internet.
        Everyone on the internet can download PERSON_A_PUBLIC_KEY.
        Everyone on the internet can use a PROGRAM to create a PERSON_A_SECRET_MESSAGE :
                ✓ put PERSON_A_PUBLIC_KEY and plaintext message to that PROGRAM.
                ✓ PROGRAM will return a PERSON_A_SECRET_MESSAGE.
        Everyone can put that PERSON_A_SECRET_MESSAGE on the internet.
        Only PERSON_A, who has PERSON_A_PRIVATE_KEY can open and read PERSON_A_SECRET_MESSAGE.


□ ────────────────────────────────────────────────────────────────────────────────────────────────── □


                [ Generate PRIVATE_KEY ]
        GENERATE_PRIVATE_KEY() -> PRIVATE_KEY

                [ Generate PUBLIC_KEY ]
        GENERATE_PUBLIC_KEY(PRIVATE_KEY) -> PUBLIC_KEY

                [ Sign data with private key ]
        SIGN(PRIVATE_KEY, UNSIGNED_DATA) -> SIGNED_DATA

                [ Encrypt data with public key ]
        ENCRYPT(PUBLIC_KEY, UNENCRYPTED_DATA) -> ENCRYPTED_DATA

                [ Decrypt data with private key ]
        DECRYPT(PRIVATE_KEY, ENCRYPTED_DATA) -> UNENCRYPTED_DATA

                [ Verify data with public key ]
        VERIFY(PUBLIC_KEY, SIGNED_DATA) -> GOOD_SIGNATURE or BAD_SIGNATURE


┗━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┛
