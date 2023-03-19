// assign a reference to the IPC servive.
const service = window.go && window.go.IPC.Service;
if (!service) {
  displayError('IPC Service not available.');
}

// container element in which to present content received from the IPC service.
const con = document.querySelector('main');

// service request parameters:
const params = {
  module: 'ergo',
  method: 'get',
};

// call the IPC service Request method, which returns a Promise:
service.Request(params).then(res => {
  // determine that response matches the request params:
  if (res.module !== params.module || res.method !== params.method) {
    displayError('Unexpected Service response!')
    return;
  }

  // determine if the service response is missing data?
  if (!(res.data && res.data.content)) {
    displayError('Missing content from Service.')
    return;
  }

  // otherwise display the the content returned by the service:
  con.innerText = res.data.content;
})

// Facilitates showing an error message within the view.
function displayError(err) {
  (typeof err === 'string') && (con.innerText = err);
}
