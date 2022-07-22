const axios = require('axios');

function getMovieTitles() {
    return new Promise((resolve, reject) => {
        setTimeout(() => {
            let url = "http://localhost:8080/hello"
            axios.get(url)
                .then(function (response) {
                    resolve(response.data)
                }).catch(err => {
                    resolve(err.response.data)
                })
        }, 1000)

    })

}

main()
async function main() {
    while (true) {
        var data = await getMovieTitles()
        console.log(data)
    }
}