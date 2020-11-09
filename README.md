# fstream

## 기능 
    
- Usage
    
      $ fstream [스트림으로 넘겨줄 파일 이름] [명령] ...

## 테스트

    ## hello 파일 만들기
    cat <<EOF > hello
    hello, world!
    EOF
    ## 테스트 cat hello 파일 읽기
    cat hello 
    ## 테스트 fstream을 실행하여 cat 명령으로 hello 파일 읽기
    fstream hello cat

결과 `hello, world!`

   
 
