# IP Bilgisi Alma Uygulaması

Bu uygulama, gelen bir HTTP isteğinde bulunan IP adresinin coğrafi konum bilgilerini almak için kullanılır. IP adresi, ipstack API kullanılarak sorgulanır ve elde edilen bilgiler JSON formatında geri döndürülür.

## Kurulum

1. Bu kodu çalıştırmak için öncelikle Go programlama dilinin yüklü olması gerekmektedir.
2. Terminal veya komut istemcisinde projenin bulunduğu dizine gidin.
3. `go run main.go` komutunu kullanarak uygulamayı başlatın.
4. Tarayıcınızdan `http://localhost:8080` adresine giderek uygulamayı test edin.

## API Anahtarı

Bu uygulama, ipstack API'sini kullanarak IP adreslerinin coğrafi konum bilgilerini alır. ipstack API'nin ücretsiz bir sürümü bulunmaktadır ve API anahtarı gerektirir. API anahtarını `main.go` dosyasında `apiKey` değişkenine atayarak değiştirebilirsiniz.

## Lisans

Bu proje MIT lisansı altında lisanslanmıştır. Daha fazla bilgi için `LICENSE` dosyasını inceleyebilirsiniz.
