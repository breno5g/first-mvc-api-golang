const onDelete = (id) => {
  const confirmed = confirm("Are you sure you want to delete this?");
  if (confirmed) {
    window.location = `/delete-product?id=${id}`;
  }
};
