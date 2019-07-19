from igramscraper.instagram import Instagram

instagram = Instagram()

# autentication supported
instagram.with_credentials('mispersepsi','warungpintar123!')
instagram.login()

medias = instagram.get_medias_by_tag('tumbuhbarengwarung', count=20)

for media in medias:
    print(media)
    print('Account info:')
    account = media.owner
    print('Id', account.identifier)
    # print('Username', account.username)
    # print('Full Name', account.full_name)
    # print('Profile Pic Url', account.get_profile_picture_url_hd())
    print('--------------------------------------------------')

