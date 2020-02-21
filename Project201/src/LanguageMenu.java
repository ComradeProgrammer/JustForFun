import java.sql.*;
import java.text.SimpleDateFormat;
import java.util.Scanner;

public class LanguageMenu {
    private Statement statement;
    private String languageName;
    private Connection connection;
    String[] functions={"学习新单词","复习今天的单词","复习困难的单词","显示学习日志"};

    public LanguageMenu(Connection con, Statement s,String languageName){
        statement=s;
        this.languageName=languageName;
        this.connection=con;
    }

    public void run(){
        while(true) {
            System.out.println("\t菜单");
            System.out.println("----------------------");
            for (int i = 0; i < functions.length; i++) {
                System.out.println(i + ": " + functions[i]);
            }
            System.out.println("----------------------");
            int choice = 0;
            while (true) {
                System.out.println("输入对应的功能的编号");
                System.out.println("输入b返回上一级菜单");
                Scanner stdin = new Scanner(System.in);
                String input=stdin.nextLine();
                if(input.equals("b")){
                    return;
                }
                try {
                    choice = Integer.parseInt(input);
                    if (choice >= 0 && choice < functions.length) {
                        break;
                    } else {
                        System.out.println("Invalid number");
                    }
                }
                catch(Exception e){
                    System.out.println("Invalid Input");

                }
            }
            Main.clearConsole();
            switch (choice) {
                case 0:
                    learnNewWord();
                    break;
                case 1:
                    reviewTodayWord();
                    break;
                case 2:
                    reviewHardWord();
                    break;
                case 3:
                    showLog();
                    break;
            }
        }
    }

    protected void learnNewWord(){
        Scanner stdin=new Scanner(System.in);
        while(true){
            Main.clearConsole();
            System.out.println("本界面中输入-1放弃该单词，输入-2返回上级菜单");
            System.out.println("新单词：");
            String word=stdin.nextLine();
            if(word.equals("-1")){
                continue;
            }
            else if(word.equals("-2")){
                return;
            }

            System.out.println("词性：");
            String partOfSpeach=stdin.nextLine();
            if(word.equals("-1")){
                break;
            }
            else if(word.equals("-2")){
                return;
            }
            System.out.println("汉语含义：");
            String meaning=stdin.nextLine();
            if(word.equals("-1")){
                break;
            }
            else if(word.equals("-2")){
                return;
            }
            System.out.println("输入注解，最多一行，可以为空");
            String note=stdin.nextLine();
            System.out.println("输入该词的等级\n0：肯定会\n1：差不多会\n" +
                    "2 不熟悉\n3 听都没听过");
            String input=stdin.nextLine();
            if(word.equals("-1")){
                break;
            }
            else if(word.equals("-2")){
                return;
            }
            int level=3;
            try {
                level = Integer.parseInt(input);
            }
            catch (Exception e){
                System.out.println("invalid input");
                break;
            }

            String sql1="insert into "+languageName+" (word,partofspeech,meaning,note,level,insertdate)"
                    +"values(?,?,?,?,?,?)";
            String date = new SimpleDateFormat("yyyy-MM-dd").format(new java.util.Date()).toString();
            try {
                PreparedStatement pstmt = connection.prepareStatement(sql1);
                pstmt.setString(1,word);
                pstmt.setString(2,partOfSpeach);
                pstmt.setString(3,meaning);
                pstmt.setString(4,note);
                pstmt.setInt(5,level);
                pstmt.setString(6,date);
                pstmt.executeUpdate();
            }
            catch(Exception e){
                System.err.println( e.getClass().getName() + ": " + e.getMessage() );
                e.printStackTrace();
                System.exit(0);
            }
        }

    }

    protected void reviewTodayWord(){

    }

    protected void reviewHardWord(){

    }

    protected void showLog(){

    }
}
