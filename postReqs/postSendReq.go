package postReqs

import (
	"io"
	"log"
	"net/http"
	"strings"
)

func PostSendRequest(trackingId string) (string, error) {

	body := strings.NewReader("scripmanager1=pnlMain%7CbtnSearch&__LASTFOCUS=&txtbSearch=" + trackingId + "&txtVoteReason=&txtVoteTel=&__EVENTTARGET=btnSearch&__EVENTARGUMENT=&__VIEWSTATE=uY8zB5aFfgh7M9dbxGfvodi%2BnsVIHA9QkJffLj0gOVgsh5BwnBOwO93BFmHJRW7rDR5p6pYIJ6H%2Bqivt4dVvRPPzfyQHme8puGbmSdu8rJNoHMZD9L5tGyHQDITjFHMaNbuk1Y6K%2BnkMHZnoGNgqWCchm3r3wllr83YufY3I%2BgwktygUXzi9AcpvGAmjBHL%2BuOPj2aPn4JDU0pe3ZnayT4gfbE%2F8x19tuFKzUcTgcEjGlrM2JTefkvjmBHWG2eUbkaO7qcw9s9SoCWDoSfFeu5a0%2FjMfsnmWyEcn%2BQgEkREaqdW9P%2FqQTCPCPo3AXzgI8o2qlk9qFdaXVTpeZxkPihlWUqcXGeNlZ9Hah45Fg9RtG3gZo6%2B4WVr6gY9rqyeyK0pobjyXI28NfH%2FGg3aPd8ewTxHUp5AXpnmIYso%2FCAr4oNxeSskNbdgfRMhxR5px4F0A2wSihT5go%2FSqgSrU2wLBhG17tjcxKnQumXTEMAXyfHGj3GTiSx7XOF%2BEoLUg42h3fYPHzPJy7NM10PngtgIAYPkNUKJ8SIrVE2%2FsJhhHgGj6MCQF2dgkeXFCbPIoaO3WMm68Hmwf8WvUt67kc1uZwtYuJ7zi%2BN8F%2Bi9ceXmEf6IoFHizrqZRjPBt6ac%2FbtaHtRlUO%2Fp1RX1TYMoc24lehmnaWCQeifvcAmRp9yyOOAwYb9Ty1SBWCtxdz0RnwvgSCqhsrKKa4JY7hPfq8%2BlTx%2F6gGyVooDz6WuxtwBJOzRlUgRmeHuYVgi6BtIOcnmLY3lSb49wfHtpz%2BmL9FVLiNbz%2FfRHwd6SGKWp7KaH1mCxwIAF3oa%2BM%2FtkhtcKzgPx3URGQ90dA6%2FG0encWbPc4advHq%2F6QdwSFLo0P1KEce2XC4M0R41QDHFijB8oywlY2qY3cCgwVxzuaDINZ4hi7m65%2BLx7wOUXifdv6c2oYrJVEu5nUW7SkBayl92rJDzlK6L6jvDzyP0hcH9OTmLv9neECdeKyBw9XEkA6GmO571VEopzMjUX2oE30dzi8pPnoxCbluiB4z%2Bw%2BEOmuUQheLMiI%2FrmYJxOI4ghUogpaSfiTykA3A68ygXA0z9%2FF80VaVRvQnRL%2B0wzuio3k3eSc6eq8ZiujfK4Bok0MpxLym44C25Jd6b65AZsOSizBuV3dFvb98fsM8crt9yeupyqQXIaOS54CccDqglV7uAUbpRSVvxoquwy65sRmWzo5naFqgK1pdCrBpd%2F25s2bqMZcXgKjB3WiCNPAYAyAd9MQRL3JEZ5pdRry7K%2Fg2MZ8btySdEW18QN7m2CkfKAOdDm7yt2V2zUvv36Mz6BRz8j0Vz2fppvMGOYiPzHjRPNY6jaJtqvHs4qWxIYmTCKH%2B4Tb1g%2BruRz65e9kBR8mYqWJvdO1v4uW1CypKXwIhtIdyJ7AeBUcbQYiaJCD4wW9ea5AJJnxSeHaQ2%2F7z%2B36ia%2BY3ZP0w2B1WyA6AajBHVVLpMmPiHwVi5lIMNPie73WL7sKhLrXeuJIrrCyk8VabOfYD237NAyKJcPJga%2FyJYvRbnOFrJ3vWR1mvXVrtNF4k8XOy2SWEcqAtHiJuRcGnEKeOMDWDU5BZThprT6GXZ9uEXdjL3f84b63WR20pewoT5AvqDezx%2B9kd53dhUTn68Y5ASkUM5XYTKlpoZi%2BHxOqmTUcTBbRK6RwOc58Sa%2BaBUXLxOVrd%2BVGiwoZY%2F%2BzJkjb%2Bj40t%2B9F80steO1Lm41GTjdTZ6SzCZMLXBkG6j1iM44BK8FFpQJe8EdsJ5DsHA5ixfrzB9SfrE93NmpFM%2B%2BEKzbPhOmhMPYHiw6b5NMIEK66hRDkjFqtNX%2F2TC3hPiI9i0Ur0oGOFR61CoywCAzuuIvHlENHXaT10oWjy9azj%2FaaRlIMiIv7jhMEy7rWk0qfhWsDjcwehIwmZlUptLcjVvEZxoLv9hjkXmv6Y3EH2zUeAptZNWNKc9dYxjNfKyl9Bp6bZX%2FBssLX4cX8I%2B81FJyVQQYSx7%2Ff6VENZnHRLdoAvQNE3gVGOPzShKZiB8fVPRBtZ436sBdmQwH%2FuDTs0Fhdm6Z%2F642TXAqIYEPb%2B%2FA34fM6dlEbdLrOy83MDYQ4CT5Qz0aI9Wh70xAHOM06xyNP09SxMjtLdPckM6yso5byyICy7F26YnseVdm1RaZgWMuyqNaUqR5DvHnsRG0MCWl%2BLYA1%2Fo77aOcbWtascTPq%2ByhN2kVEdfReT3rjsX06GiIVSi1mdRy5o9kK4FE2%2F%2BNHleddohDpnNOEkM24m1kwfJ561kOuZ%2FX4UTe6wWihqQideao59rPGXraGhOVs9IjInSK47MYFv2hfY5adMA9XREjdhtPR10qgoQQqzAGX16kpT14PQDR8tRYki%2BjypXPiGqqr7xHBiTvoCEDm5N8nc8nyPTiQJRYeoleZexYMXCjQ9BnAqFDYhFiG0eR4sW0osYV%2FXm%2F5xxQx%2FjWi8DEFVH0mPsaDaV2LI0NhL9Pla1NsGujvQqZTvuAcITfCGUdtkmSFJS3axwrkjElJjBue%2B3maaEyRTmVndYDwNnSUZKLe40rYmxPmlQ3QpL5BKELetlO8cM1xQMGAuy5o3QbTVjpABia2ARilTYLYyTlBYWiKbVFFDiU7AnxzoRIiEIiJGFWrwnug17Vj5eOSKs%2Fw%2FA5nlX10sjVATFqULE4W6NnIXH6%2BBGjKYvBzuWoqMGh8cYxcXMz%2FGXvZ28h5SLUdZK6FHbApmr0QxO2UmMQcuEEzdPVUYM6aTzcxhAPjYJswNpuyL%2BeZQG3ERdaA8BqZe3Db0%2FetRU4he0uJgiDtVI9IESdIthBCsH%2F3R%2B9mhrfZfoyhENA966NDT%2F4zM%2FkDvDhWEv4upTRa4ZxJxRVpXFJq72a26z412AyakzH9xLpaf0GaapnD%2FD4l0QtnhED2M045yLWGG1%2BEiU3rsU%2BZQ5JKRnRPicKFIDH2wjyUhKdvvYapqIgYDlcr%2FLRloIgFdXdY5IdxGpJZJd8rUV2VQ3MxMhw89d2VQ3RaAh3AC8BAd%2BZUc0iQg6%2FJSDmJYbj7mq7nXIlI%2B3Gb8eOIxwoOv8wh5xSc4jpZtmi1YZ8NuGECYLN60bE4BRYoe0jeLrgEDFMhowJgKuEOvtxl3rtG1RtctnezjK2Dhav%2BkEstiQigC8mSNR2fJbz1qYKDvdvEMaSboPuFmfzVLLu3%2Ban3pY3lYUrWHqWrH4ZbCu%2BhWp4eYFFalUY%2FUcYbe%2BZ12SszFDfI5lV%2BI%2BHPb%2F2w%2FB57FOS8%2FlGTe0V1QhL6tjDQXb8BZRBbSufX60gxOLaKWs8xC5cP75DhI4IA0gl74T21SoCuqocsXm6ZZd4tLZuTkg%2FE6RXoXyytZ3kdKYRX8ahzBAI0M%2F8fKjGyGelV93Zgf%2BlydgAoyqfufhsXvk2DfE1Xgie2hBG9bOssU%2FfuZq831Xve22rwneH5QNJpeXCh9rKQpjJqYVs%2By0Bd2aG9VmBCGUWunKbMr29t%2FHAHsf4c380VNKfN7%2BhyDhqZoZd0jCtmjZZF6h9I3uiS1Gx3gvOMCazcw%2FaqUG%2Bf7wKnd0HNZyerVoYo33LWYUhSedEt0xPmQNRcRFIROFibYAaLAH%2F3b6LuKVURTF0EqElcp7QCCplZupoQJ%2BUE5v04Afp5eVC%2B%2BA3HzhVf32YpUA3VBhUP3uoVcvlz7A%2FuikFGXgp%2BugGcampH%2FShYSxIerFVoV9nnl7Ok30%2FSoPENM8S%2BrqDab4RD03AJo2h8Kc1gUiLiyXELhwBkCqJqqyyWZh3trO25M9j%2FDGrRq4FsW4beIXp9hhIA9uyc%2FAQSGMR4NNEW%2BC9%2F8OY2%2BmDlzRTMMPstgHl1XcmBz1ha9hkN5rXfb8hvQ%2FK3wyoz0clZgdwPLJMxH0b080uUC7EibLm0jALzHnqXEGrQDHxaT0wRQ3mRPT8VweD2mpNrmUNNQg5i%2F148qHBJfOZFRLCKeb%2BWJkwvxHSLRLexrQjfR5UWsieJu5RJImFu2RrC8a%2FP5xUnr52kNV6tXeZ1cTkaWnO2UWHzDc%2B%2BnAKDzUXOH4mAEXnD542b6p69kHmmxOdaRibzrwukPGtvo33FVuf2vmAazzxwqLaIYY1PZHNP6XLmk6H6wjzgpl9PAJmDfz0DnfsJ0MeS6AQUR1vN8t8HJaRDW6U92XfL67RXiWgTBXP3bwio60FiaKm5Zri%2FrEVmHtvmwm36u2sPocWAQv100%2BfpjAW3DYTD52mAu3v%2BHzlSuAH7L0dWb8uvOh3xO8CMobW%2BsGhAJ4tbrl%2B%2Fi%2FcUhDOhxXG9uZ%2BpnGBEPEWiQzi7Hh1i%2BUmYVuEHcKqnVsnHvKMV8IPaspReWQQ8vWG99xDu9F9rPOLLl6Vw8j%2F0XjvrkU6NoRh%2FkNcMdsGYm1D08KXCo380mvfxdM67w6A7NH8EORRNpwCT1dHs4k54xKOSMs4HiyHGcKMYcgL%2FGwyLlmDvZrh3RoTFo7SM99RkhjoyuhrVfv5g7IYkk3uEyhpMbnm310%2BuvBz6n0ESegS2UEvJJm59xNWh4%2BLHtOckOvu3mpGUg95FuvGFwCFfY8ycZlmPURf6w5txdxUiZhehHTzrNXYDwLC%2BIzLqnI1KyoC0tB2rx%2FB94Zl47HhRXEDonbO04CJDO%2FPOCHv14L9wovacT5RdSZgHYVNQuz%2Fm3uV%2BwfXRIZTj8tW8orb2T7m9mLUwhjVvN3EHm0dFeostbnV5r51I0BRi0mv82PBVx%2FORZotpEA7ouVusqAnYV1%2FGyQVO6j2R6zvlug22jEZ4PkCOZ48byJ98PDcry6uWPvJ8V7pwjl2YlbFVimLVJDCrFn6TO4VIqZGtuv%2BI7M%2FPMzPCpKUsxk%2FXlDHBQzKXzh0R6fCFJnWCDsVkmK87b6rWT9trI%2Br5ZOvNhj6RpNx9M4Jkn3ka6jAekzivQJ44Y68zr5ykY6WRhsojfKHzYEgD%2FcuQ78yuy5N6FdHtt%2FGqKIPdNVA3kJFVPsDInqZYtRmU1Kn93TVElOXARFS31XOtW5sfocyQ8m2p7LxEOJjVnIGKLJsEH9YLD8Pfi15cAjY2bvqqG7gJBOYWbw2i220CYKZ9fMxO%2BHocX2%2F5ociCpplo6P0ekDcXUn7vhuOPCnUbIkYJjSyOgP0L%2FS7k%2FsRPbmR%2BopeBo6g8zlv1D8PNRogcRsVmfwG2285YaoGvR%2BqyHAcTNvmIwT64Qi546lL7LgRaaXdygjQH9fyTXNyPoeS8ELzcSpYQWBT%2FOIufOfD%2BZdbfOk9Uuw9WNzVc4Muliq%2FmmnEfFHLjzIBNOabX16EEFTmlW48O%2F%2B%2F9sr6UfSwy%2B91orVbIXrueTTDbwQjLHTfNvtRxj4b20IUc2tr3QQEl%2FUSpOmVfS3B6jKQ4Ho46t9gVY8PfrhENkzNBb98kgsd%2BRUJwmCcObR2F3bfJwVxuj3PKv1LSHD8k9FEBReJqU%2B9en5r2XS6QS3eMLsfTSyunKkj9X%2BfhE4ecQrsJWrauk%2BmP98IA6A9CidN8%2BSHa%2Fa9LWEhsbdJ206mhPylUBzes9gCuIuqyARaoGQg85i%2B3XHp5%2FmZjOA%2Fr%2BaxhIQhloF2MmjECsMIj1MWzqUidbc261Le0Vahq%2BIQSajeIUVtbkcNW%2BR8g94OoJdjcjbmRgyYtuHZMg0r01O0eSH2r39QNXJG2KsqtkZKyMZxfpvg34S%2FWaGTDEfEudyEt3Sta8PXyLaxCOHIfQuZWyQgmpavwkTgeEmW%2FTW48oQkqUqueYia2IXCqEI3pWhiILgdFAA1rtMyJ6tcVzoQoqbRw9AFAl5i0QHlcnXfY0xHffa3nPULIAPP3Fu2354tUPh4PC0qftV6R02oiD69gKDrL68CHGzJsoMBt6BMUvq3lfKmypIUx4m2bpq7kav%2BWeMVkLJomFInufbvgk3UK%2BUvyaszfC%2B0utWws2AvgABPL7WJlp1x1Sg%2Br6I7XSQsZ0fcenP4ZLhSr51O5Joj2XE77c%2FmfF%2FKiVuhXecjLkz1OQVhBAuEap5w6eQlxI53zzXW5z2pgyWZhXZOMcE2esNtWXgnzTuVW3ZmcJ73SwchplzbOegKeMniWpCkd%2FAzQzMShKvhhPOj5gXrV%2FxiigUQFHeKLive1bHYemK7tJ%2F5TVNF9CFykdVxOvDobAz%2FNFI%2BLg99U2jHZPIBwzgij3FYfL%2F7IzxqDeaOPMr5E%2BdhZSaBmcpol0a9NQ8MQ5zTC1vjgb9%2F0e5%2F2%2B%2FnqgJyOQwrPH%2Fh3jYbYpI5%2FG%2B%2BQUlUM7RIpqWwxOWDCehMnOAR8xh0D2m8cD57oKKcZ0wHnLafkLfFffVlalTm0FEdaKdPOmz%2Bl86bTkwXJyF9ahzD9EOMFeR%2B796Z4WQDExv9EvKBbavUSNpRyBfO4Wowd4GoQzTVHs62UfHr1ZIRQL1saIhv6O3yiD4q%2BlnI2t0bB0OQNetPPz3rHbc1uj1ZAi13EaCFSoSaeQsp39w5hwrnVkL35RQ0cLtcA%2Fu24YLY%2FYP0rz%2Fgmdu3pUIpzKZcnn2C4eyvGna85HtuuycEcEyA%2B9LBKU%2F8FP4wiOF3VXaXPoB0HxCHDBCeQKusl5H%2BJdPipFjbhwx4hd2rHy6t1IUYCgVCugNb9YnOEoZfAU49FlzuYaM4zONm5PaqWWoWxWv8J27wn2kdzlc2hzJ7YTBU7ZRMVcZNEGUyLWwosF0GksJpfzO9EeyOb3Pk8Jrv8dtwoR%2FZ4%2FY%2FhMFMrkJ%2BoLqgmlxZvSe%2F4IhutVbKvC6obi7sSLvHfuzZ1U9du%2BuT25lGWiYXFstUwcCPVhVJ9msQ63dvFmO0f8%2BHqQUUuJi6ol6GWt227ycCdXw9Wae%2BjOZpR7vjBuzh74w%2BQCuFOFG1ViYJAXiFrW8waG8hXh87mXdfAmN6OFbZ5aGlGi5ar9X9BaVu5PxSaK%2FAjVr4jd0%2FvCYhGm8GA6cVFt%2B3pwimU8YKguS6H6I%2BcEfWdB%2FpXw53pRaOATp%2FZ20UNSSu%2BMc3ZSIJFdavtIiF7u5B2tX%2B06VAZWWv9T8Tr%2FnSrPcfr7dZGLfb05YmvvF5KQ7raBs%2BlzUkOWrbE6OMPGuWy62Ik%2FiX0g1FO2V296E0chydUxIar7JPB%2FmYqmoM27VTaLzHSS28WrdtXQPX%2F1uzJUYEv78AFevHrwh%2FF733rEdo0NObwjsDJB8yYjcDNDC97bb5T7FFL0RGiV3tCihlTpNlfQu3HcaDHMbY3sbPCEUg8lZkXLnEhTfpIjn14RLoiEba8X4j95KKIGAYcoR2HDNBiUeaUVNS4YHqJdqc0JucxgolH6xwsMEa5o7CmkWaGWUfs%2Ba61xSL9pE4avEKaZsiqA%2FTvKi71tuUcQbj0olsopQl15qDOiQ2lj5A4vlZGSi8LxYenv86sNT8U16yC0R2Fq9olf3i0AMYTjUrg0sUKmsEWkn5fNfoPnBj6XkBkFn%2FSyH9bVXWfxSIi0xuXS88LR2QSityNgDJcZEzQTPllZw5IeqnMdH2i7iW5crZ44WlEK3OUhHbOXW77f6T%2BIkIDSTbK0KCnPtdH2QdOBoghjdoDPj%2F0N31hNgzCbzBweKyzXPMH8RxIkj0Mwvy8Ylfop2XuDaQYwfBaEXNKCvFGAB5lKcqzaGZNfHLy%2FLEi2j8iQPnod7LnfEHQyM9rz6hb8PV0wsoNK%2BGQHbLcmvy5l7cuLGWJKThB60c1kA3%2F9cNpk%2FkFfHGSohrmGIxutYU%2FHG3WMQOi1uPDfCM1HTcRrOeGXZMPwHvQ846qmUsucO5qLpRugN98P2jVYotAoh%2B6Bm34GqopP6xE9A1xzJlumUIZMop0QMQn7ukWLTin3sSzG2BXGYCzDEnfeXHmrVQh3Vw8Ty0gLHVevVd5HCSyh0vGStnAtLKACgM3GcXQyNUmD4nNhw1lbM4ZH5Lk7QG9jznILdxrR6Je2wjxSqQOMwNb3tXZJAqa6rZWxJaeWPJqndzM8BpvZ4x0WoVglYi23HuOMDgIwWfaez7n9PHom4%2F7cZq%2BGN5Ru2R6GwHBainOyuP%2FmyIJS68T4caFfRWVFVztvw5cAF0hlSk9weNw0CIfsFcfSIK8upCBlWt6GdP%2Fr57PNeJBbfjz18NG2qFFQodzWYTPto%2FApVBlsb7koC5KGBxyRnSq0GY%2BP8l8vRl92XHL%2FYkboSTtLug%2FyqrW0fsWpF9A8Ugvfp32456PZijbLK818xGnfFm6kZ4QosVUb6DXJ8A5Thzy27b7eG8KUi09p4ltpNwYBSOQFs%2F7H19b99YvFQqmjmDcnFmzJwrlRjC5WDDVTb3sqcNf5P340ARc8iX%2FdssQmY6VNcEZz2NpUp7esPMA4VgYfWbE7wDejSBkMe88eMNrTWDDI90dR3BphKqiqBYgnvk0B2G%2BWYSthztHHhorB0dhOJ7PHVY0kUmOinGermXvUJT15z%2F1envGRW8IrLLiA2OnoF0EbSLHf8qwxtt127KlcNZmxorGfO%2FgppX96SRtXgBIhxfkV2Bs2Xp3395OEo6TstOeJqmJosTjxTLLGwdkXhKB6E8O%2Br2ZClflatAgLOAjzYvHgWQpvi3YlBQ5290rqkXssDkv3pxlhI1CbbwcD%2B8IjoYwI2l%2F6ob2yylFYdB4koe3svVqqokS8kjcPXGVT2oXEvsudEn3Yi4yLqLHQkyFaWNN11nP%2BNKRSozMWij4VRJkNfVlH%2Bj6VYEyoecAhr6wzz9I8wPt9JUVLl5PL1XYxDvLPokoFCRvy7f%2Fc1k%2FjBYX3CDxP2%2BSf%2FbEyRj0pZCXda6%2FUEG9%2Fu8jNs01QrGKWTRFsjfFV4UNwB4sKgeDiL3qTXhy0SApQHm8oftBKRjKsmVh5JtfjEK%2Fv3Fx4Lw4LuXJPenVzCQjIoxq%2F4kmXqbJex5iVZQRS22HtAEGHLK41Jz4tzb6LuD6vqZKjEVJ8dyePBGXnMjhMkFmW2ZHBm%2FFhCoy6k3Z3Ag5fdN3696H4eYsmsoIOJVY6VmchqVZuJqO40DYTLCVQ48hzjOM7Nu0uMrQ6uhEs6pz5Lln9Mx0XNQXL15KsKlcwvgX0qf6q9LWzZXXg2COXSMx7eXCneZYazOxC%2Fcsq5BMw9i9HMcyQtArGXMvUCSGekO6FTnOUa3YySPSxZ9IKgwAjJWfrrWRqcFAMkeJ5qpBl%2BZaEk7YxO%2FPtXsh3N1HpIjelRnrIoWlHsS5sonsNoXrXiQglFVVfhbRIYWTOk%2FLuG1%2Fz33W2b57tAO%2FoqhSP%2FmpexT5NwneupiHMf0OjHjiNR3pgjDXxkIGd3i4ZjDG8bAZFKSwYU%2BWflYTN3ixLEdDIXgUPceI53cqtrX%2BShkjf0KeIc7eTigLbUppaRHaTpc0yWSY0wWSNLfpeaBhyhNd4WCIOmxSkKaEypK2IEgjOAj%2BqxYVvEWW6CA%2BrFXPeV7KM7CYYkODWmdux%2FxsHN9HHKgdyBNiffk5%2BnfosAgpfbUK4r7W%2BymgPhIDYFaLfI51r8063cbo%2FeadcfQHiTfjQVu8bvEhBw%2FIS7utBnW52k%2BoAHEE3yXlXkuJQbMGBMVWTnXjfB8ahDmrgRhvPyJ%2F3NC%2BVsEQHxzswPXKwWd%2F2Mpc%2FvtU9ak2JztA8ROxifplNQ5z7DIyabWKRl8AzYx1VndLldva3pbNJwleNVNxP%2FBpkQiPzSKbZVJJz0%2B3hbTLujnsDlS9%2FxI0tAi%2FlD8LxCxzE8Rqum2NeYRsRin9rVl1js4c26YqO4t41TBiyXNPyGN0Clzdyj8AZYXmnWuqI6YheHR2TSbSsI7q1atwvlMcSaE1%2FhSzwuHbQsVKk94skQgQ1kTG%2FuFnsa%2BxPjBNlF8mVAesS%2FT4%2BZ8%2F%2FCE%2FuB8gCkRLM6U%2F6ECcqf3%2F8XrOU8Kt9sqHv1kGs2ft8NUjD%2FMLhy8MFd7rWeXRaKuXIdou1t35ALuWZncyJpL3MSAmsLGKU4vOXwYO810A%2FI%2BDCYxEgglOWW%2FJQnyUtZkb5iz4qcFBnt67q%2FNdbOSm7zJUGe1Pz7RAZvfBRynM%2FadBu8Ip%2Fg2eSjeT8OCupn59JC0xB0s3VDaAfhIDP%2B9xrcJLMlAIWPl5YZfT%2F35nVnoBm1d3S8Pi9cz05qcwMD5j1VR9IRivA4qQrCzgX2KWFItC71K0ZI9DoJN3WjcASaVXL%2BtRYMghqbXIttzKJ0yqca0hYi3EHPCgS2JQkG0XW%2FjgcjrqX%2FBRWZSJXpD05AMaeQBauN1V%2BS%2F9gQvIoiRfPtR567BRM14rh5aJPaNSLXeY5Xlwmm0SLdaVUwLu3EW%2BQmByx%2Bqzkjcd0Y0An%2FB2Cl36atvrr7ifSWJT38HDjst491iJV9vy5eR8%2Bs11VIX%2BjlHnEJmrvfvabC67DVnc8KPmr5qZKP6bMz7cdr7wkg5VzCl7oPB%2FeCWdLwx%2FXjpCVWCqnFYBUvnDtlXvt1kH78f6m7TvjwPN7J7PMexbJD6O855h%2FYJVoy3V%2FjWQ7pQUGlTwDrxE2H9huAtEpoP6ZPM5d%2BmcJb4EgKV%2B4x51abi9VMgL5K6%2FVuXpaz7JqZmZ3Qqtze%2FGtYI%2BLZG1Ax72JbW8aDNYCxX2ArHENdbBWF%2FFjkqOT9stTq1HNmx99J%2FUMFk4hJVbGFPM9FSFLuvCnMRxb8psP%2FPhYVQH%2Bg5yk5nxmwH%2Bap5jm46FVTyRKBLccinsBCeZKU72dI%2FrIPBFBKbsjU2RguqxmVH5a0LkTCX0zEPnco9Wdiod9FCWgt8CvRdR7Oflr79gedB7WdJjeIYuTyuM73rBtKKEO4qxI7qcG13LLr55OOSkYoUy8X6T77ffjjY%2FVZXgoAEyWco7vR271%2BZSrac9U0uT%2F52VBTQXbiCukEwi2b6qBq3UIpzlkIkHfZwTVh4opmJV3v58KFP5KUznKI04OvQsYejhZyfwHrjxwPQeflSQF7jX66Pn1yosMeijDHyAZpejTKIHbZgPf8cIIbtOPd5yzKmyUxiyZJYv5%2B28bKt4C7dGIZVJlFfvhZuvmYK%2FIEgeKnh8LJTATlrcHkpVGaa5dvuUW7sdeZ4iY5rKLVVCmxztzA33m2b%2BcbajT%2FmYqJ%2FNa1weIq8CgmqdiCemr6d%2FrKEeScinpP%2BWH5voOPvn3QE3jEIMQpQVHvk9Pmt21SUYe6iEZCbBiuBv414jm7l%2BM5oDDpt8cSvlIPhdEXlB9S8pcKCcSWZBT%2FWlVx4owfp%2Fazs4nyXg1u2Yeb93fdKCQXHlMmxx0LdDGKQZ99bHvfY2nERg4gmZGEBSP4mh5hgBF9IbC5tlCD7GzVZRItoyDfOa9CtS5u%2BFmUmVBtP9m%2FuEfVYbUNxWYQQ8U%2FgkhAXZP6vQfdU7EFpcihdnWGQHxKI%2BWlB3nDmYF0hnrZMdDzyLFAGwnOL0fyQ%2BF4bhfOpwOp9pR3JXcBKkUpVS66F8GqVQqGb2QlHSB%2Fk1kGrsK6wFvdZOp4W7BGpWvlFWE9ABKjiBOEwHjTB21zPAPWnWIqSHYrZkhDVigXkKMq050r9Xnyl6vMudVzc%2BL0HVKcFovNfggjycnsQGfJX%2FKSMXGzy%2B4%2FXLJwopgAV5Ffec7xnJGj70wjj0Hsw1%2Bl7sd%2FBb%2Bjjk4LkQ5UGWDjRxnALi5Req%2Bfp3uRxO1ErhKp1Cbw%2FJjzIWX5ciVxOHmECoosTQcv%2Bnfa7vBdKLHRPggUMCotf7U3cVK%2FZDBOzXNVJ7hXsm0f2%2FbN6QYdPaSUa35XchqbqyLElNkOaE51nlQYN8y4vTVG5G%2FwUcjCoObbpl%2B407FbQQ1%2FCVQ0sXih2Yk5Y7SHXsn1yuosPeUb1AsjKtYOustG5caJE5PhPNpd7fzEw36cXYeOjIofoYMem7CkbtNK1pdSkMPZ%2FOMk2M2nnF%2FTfmCXoGvpA%2FFzuzVhGsypngrnFPXP2Lp%2FHv2OO34p%2BYfFXFvesbrzarnXBIeojpcl0BK%2BGZJ3u6a0woHy0yCiLnmBUros7Lttf5RWi8vsTUD45x%2BBZMv5XBnpO3rVIFZDO1HI2zOsXuhFIGapBg91BYPoSglaKOSvfmCJQH4j%2FegyhwYZQgCKdnR3tN0rADlOCBX5Birh3UExhH3OghR4MAf53ddmHiuHMTtIK%2F5Ac213zzAHQG7VudCLrS0EJvbc5bN%2FoOHISIU2c9q0332KeXSDN39rxFSpM8wcAQamYgbC3lFpUi2J8OB7Dkc4jPXJZIJu3deGGHj75OxQy2nuNDFB90ml07MB6hdqVTMrpypFoaVOZniYeh24r485%2BCYHXolrb6AXgTbivNklYRc2KE2nWjiOx16RvwsSyQ7w628tiyVQHd0xwP2rYAvhYWWKtTws6GTV%2BCEWB%2FakFmMNdH8bCkh%2FXyOt2pNo05owH4DjpMuw%2FNj5KOq50uxhVHaOF97TIzfUwte5Puy2r%2FuPPyHEHbYhFxDl23Y91cDrgCexOY2dvp1G4PwhT6lt8eX96OnZFWXefH6x45t9MmxWzKAlbHwNOW9yUPEE4UL0xAPu2dtEIdpZMELDEjGtQw%2F%2Bd8%2B7A3z63ZVPWfOwGefb6iXh%2Bc0Yk56aXKwgPAOBmjUih%2B6sjO2bei3CiGZDTmHbDKSesIZUdPfmqINFLjB6gBe9cGLNCyDcbZMlWDQB8irF49fhtfdeDZKaLHRkR2UMKVYqUHvgf05lDiA%2BAS8lRaI526bsMoo6MLevtIMQ2PWWVv6HJTis96lT5huS2%2FoeNj%2Fy3432L80xWlRvYT1EqspLFSbeCjN%2BYc2HlqY8g%3D&__VIEWSTATEGENERATOR=BBBC20B8&__EVENTVALIDATION=Lew%2FX2rShXD0M6UwTuub4nMlrIoz4jHMuiubd9Vqo7YGYy%2BPjRbbVTyomtp6ffsBvPYOHBborfDGlBceogn6ty4KkEMQ%2BAAxJrU9drF2HeJe4%2FWi4nVVBkIR3OfiKGKTSfCOrWGwLgf8oEoNttT1viIg6C3aOxN2cnTLJtBfOFFg96PGcotjo8SNcNBVrIXIubsY1zI2lvAJ4qpnRDHXKA%3D%3D&__VIEWSTATEENCRYPTED=&__ASYNCPOST=true&")

	req, err := http.NewRequest("POST", "https://tracking.post.ir/", body)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Accept", "*/*")
	req.Header.Set("Accept-Language", "en-US,en;q=0.9,fa-IR;q=0.8,fa;q=0.7")
	req.Header.Set("Cache-Control", "no-cache")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")
	req.Header.Set("Cookie", "BIGipServerPool_Trust_24=402718986.20480.0000; ASP.NET_SessionId=tpu3w4co1zrstifqnslsqygv")
	req.Header.Set("Origin", "https://tracking.post.ir")
	req.Header.Set("Referer", "https://tracking.post.ir/")
	req.Header.Set("Sec-Fetch-Dest", "empty")
	req.Header.Set("Sec-Fetch-Mode", "cors")
	req.Header.Set("Sec-Fetch-Site", "same-origin")
	req.Header.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/110.0.0.0 Safari/537.36")
	req.Header.Set("X-Microsoftajax", "Delta=true")
	req.Header.Set("X-Requested-With", "XMLHttpRequest")
	req.Header.Set("Sec-Ch-Ua", "\"Chromium\";v=\"110\", \"Not A(Brand\";v=\"24\", \"Google Chrome\";v=\"110\"")
	req.Header.Set("Sec-Ch-Ua-Mobile", "?0")
	req.Header.Set("Sec-Ch-Ua-Platform", "\"Linux\"")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	bodyString, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	return string(bodyString), err
}