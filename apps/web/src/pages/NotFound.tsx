import { Link } from "react-router-dom";

export function NotFound() {
  return (
    <div className="empty">
      <h1 style={{ marginBottom: 8 }}>Page not found</h1>
      <p>That page does not exist, or the link has expired.</p>
      <p style={{ marginTop: 16 }}>
        <Link className="btn" to="/">
          Go to the home page
        </Link>
      </p>
    </div>
  );
}
