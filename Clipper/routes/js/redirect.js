function uber()
{
    console.log("Inside Uber")
    // var payload = {
    //     name :  document.getElementById('inputUsername').value,
    //     email : document.getElementById('inputEmail').value,
    //     mobileno: document.getElementById('telephone').value,
    //     message: document.getElementById('message').value
    // }
    // console.log(payload);
    // axios.post('http://localhost:3003/texiservice/texi',payload)
    //     .then(function(response){
    //         if(response.status === 200 ) {
    //             console.log("Inside send mail client");
    //         };
    //         throw new Error('Request failed.');
    //         //console.log("In client js");
    //     })

    window.location = "/texiservice?No="+"2";
}