export function deleteUserOnClose(){

    $effect(() => {
        const handleUnload = () => {
            fetch('http://loginpage.local:3000/getusers', {
                method: 'DELETE',
                keepalive: true,
                headers: {
                    'Content-Type': 'application/json'
                }
            });
        }

        window.addEventListener('pagehide', handleUnload);

        return () => {
            window.removeEventListener('pagehide', handleUnload);
        };
    })
};
