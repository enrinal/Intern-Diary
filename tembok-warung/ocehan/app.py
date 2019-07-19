from igramscraper.instagram import Instagram

instagram = Instagram()

# autentication supported
instagram.with_credentials('mispersepsi','warungpintar123!')
instagram.login()

medias = instagram.get_medias_by_tag('tumbuhbarengwarung', count=5)

for media in medias:
    print("Media Info :\n")
    print('Create at : ', media.created_time)
    print('Caption : ', media.caption)
    print('Link Image : ', media.image_high_resolution_url)
    account = media.owner
    user = instagram.get_account_by_id(account.identifier)
    print("\nAccount Info : \n")
    print("Username : ", user.username)
    print("Profile Pict : ",user.profile_pic_url_hd)
    print('--------------------------------------------------')

