import sqlite3


class Open_DB:
    """
    The context manager to manage the connection to the Database.
    """

    def __init__(self, name):
        """
        Sets the initial setting need for a db Connection.

        """
        self.name = name
        self.conn = None
        self.cursor = None

    def __enter__(self):
        """
        Used to set up the connection to the Database.
        """
        self.conn = sqlite3.connect(self.name)
        self.cursor = self.conn.cursor()
        return self.cursor

    def __exit__(self, exc_type, exc_value, exc_traceback):
        """
        Used to taredown the connection to the database.
        """
        self.conn.commit()
        self.cursor.close()
        self.conn.close()


with Open_DB("db_name") as db:
    db.execute(
        f"SELECT MAX(year) AS last_year FROM table_name")
    return_value = db.fetchone()[0]
