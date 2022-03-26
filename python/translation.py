import requests, sys
from bs4 import BeautifulSoup

# read terminal options
args = sys.argv

# headers to avoid IP block
headers = {
    "Accept": "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9",
    "Accept-Encoding": "gzip, deflate, br",
    "Accept-Language": "en-US;q=0.8,en;q=0.7",
    "Dnt": "1",
    "Host": "context.reverso.net",
    "Sec-Ch-Ua": "\" Not A;Brand\";v=\"99\", \"Chromium\";v=\"99\", \"Google Chrome\";v=\"99\"",
    "Sec-Ch-Ua-Mobile": "?0",
    "Sec-Ch-Ua-Platform": "\"Windows\"",
    "Sec-Fetch-Dest": "document",
    "Sec-Fetch-Mode": "navigate",
    "Sec-Fetch-Site": "cross-site",
    "Sec-Fetch-User": "?1",
    "Upgrade-Insecure-Requests": "1",
    "User-Agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/99.0.4844.51 Safari/537.36",
}

# to accelerate testing
s = requests.session()


# check if connection is up
def check_connection():
    global headers

    r = requests.get("https://context.reverso.net", headers=headers)
    status_code = r.status_code == 200
    if status_code:
        return True
    else:
        return False


# generate the base url (without the word)
def get_base_url(language_origin, language_target):
    return f"https://context.reverso.net/translation/{language_origin}-{language_target}/"


# return a list of translations and phrases
def get_translations(word, language_origin, language_target):
    global headers
    global s
    base_url = get_base_url(language_origin, language_target)
    url = base_url + word
    translations = []
    phrases = []

    r = s.get(url, headers=headers)
    soup = BeautifulSoup(r.content, "html.parser")
    words_links = soup.select("div#translations-content.wide-container .translation")
    for link in words_links:
        if link.text:
            translations.append(link.text.strip())
    phrases_span_origin_soup = soup.select("div.src span.text")
    phrases_span_target_soup = soup.select("div.trg span.text")
    phrases_span_origin_list = []
    phrases_span_target_list = []
    for link in phrases_span_origin_soup:
        if link is not None:
            phrases_span_origin_list.append(link.text.strip())
    for link in phrases_span_target_soup:
        if link is not None:
            phrases_span_target_list.append(link.text.strip())
    phrases_zip = list(zip(phrases_span_origin_list, phrases_span_target_list))
    for i in phrases_zip:
        for j in i:
            phrases.append(j)

    return translations, phrases


# list of languages to be translated to
languages_target_chosen = []

# dictionary with all the available languages
number_to_language = {1: "Arabic",
                      2: "German",
                      3: "English",
                      4: "Spanish",
                      5: "French",
                      6: "Hebrew",
                      7: "Japanese",
                      8: "Dutch",
                      9: "Polish",
                      10: "Portuguese",
                      11: "Romanian",
                      12: "Russian",
                      13: "Turkish"}

# start program
# print("""Hello, you're welcome to the translator. Translator supports:
# 1. Arabic
# 2. German
# 3. English
# 4. Spanish
# 5. French
# 6. Hebrew
# 7. Japanese
# 8. Dutch
# 9. Polish
# 10. Portuguese
# 11. Romanian
# 12. Russian
# 13. Turkish
# Type the number of your language:""")

# language_origin = number_to_language[int(input())].lower()


# print("Type the number of a language you want to translate to or '0' to translate to all languages:")
# language_number_or_all = int(input())
# if language_number_or_all == 0:
#     languages_target_chosen = [value.lower() for value in number_to_language.values()]
# else:
#     languages_target_chosen = [number_to_language[language_number_or_all].lower()]
# print("Type the word you want to translate:")
#
# word = input()

# start main program
# get terminal input
language_origin = args[1]
language_number_or_all = args[2]
word = args[3]


# check for word not available
def not_found(word, language_origin):
    base_url = get_base_url(language_origin, "spanish")
    url = base_url + word
    r = s.get(url, headers=headers)
    soup = BeautifulSoup(r.content, "html.parser")
    message = soup.select("span.wide-container.message")
    if not message:
        return False
    for element in message:
        if "not found in Context" in element.text:
            return True
    return False


# check language not available
if language_origin.capitalize() not in number_to_language.values():
    print(f"Sorry, the program doesn't support {language_origin}")
    exit()
elif language_number_or_all != "all" and language_number_or_all.capitalize() not in number_to_language.values():
    print(f"Sorry, the program doesn't support {language_number_or_all}")
    exit()
# check for word not available
elif not_found(word, language_origin):
    print(f"Sorry, unable to find {word}")
    exit()
# generate list of languages to translate to
if language_number_or_all == "all":
    languages_target_chosen = [value.lower() for value in number_to_language.values()]
else:
    languages_target_chosen = [language_number_or_all]

# check connection
if not check_connection():
    print("Something wrong with your internet connection")
    exit()

print()
# open file to log
with open(f"{word}.txt", "w", encoding="utf-8") as file:
    # does the translation for every previously selected language
    for language_target in languages_target_chosen:

        translations, phrases = get_translations(word, language_origin, language_target)
        # if language not supported, continue
        if not translations or not phrases:
            continue
        print(f"{language_target.capitalize()} Translations")
        print(f"{language_target.capitalize()} Translations", file=file)

        # if one language selected
        if language_number_or_all != 0:
            for translation in translations[:min(5, len(translations))]:
                print(translation)
                print(translation, file=file)

            print()
            print("", file=file)

            print(f"{language_target.capitalize()} Examples")
            print(f"{language_target.capitalize()} Examples", file=file)

            for i in range(0, min(len(phrases), 9), 2):
                print(phrases[i])
                print(phrases[i], file=file)

                print(phrases[i + 1])
                print(phrases[i + 1], file=file)

                print()
                print("", file=file)

        # if all languages selected
        else:
            for translation in translations[:min(1, len(translations))]:
                print(translation)
                print(translation, file=file)
            print()
            print("", file=file)
            print(f"{language_target.capitalize()} Examples")
            print(f"{language_target.capitalize()} Examples", file=file)

            print(phrases[0])
            print(phrases[0], file=file)

            print()
            print("", file=file)
