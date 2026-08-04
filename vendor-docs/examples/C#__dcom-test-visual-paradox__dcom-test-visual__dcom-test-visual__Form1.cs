using System;
using System.Collections.Generic;
using System.ComponentModel;
using System.Data;
using System.Diagnostics;						 
using System.Drawing;
using System.Linq;
using System.Runtime.InteropServices;									 
using System.Text;
using System.Threading.Tasks;
using System.Windows.Forms;

namespace dcom_test_visual
{
    public partial class Form1 : Form
    {
        DocImpServer.DocImpObject o;

        void InitDocImpServer()
        {
            if (o == null)
            { 
                o = new DocImpServer.DocImpObject();
                LogOutput("DocImpServer instantiat");
            }
        }
		
		void ReleaseDocImpServer() {
			Marshal.ReleaseComObject(o);
            o = null;
		}

        public Form1()
        {
            InitializeComponent();
        }

        private void label1_Click(object sender, EventArgs e)
        {

        }

        void LogOutput(string text)
        {
            textBox13.AppendText(text);
            textBox13.AppendText(Environment.NewLine);
        }

        void LogErrors(string text)
        {
            dynamic le = o.GetListaErori();
            if (le != null)
            {
                textBox13.AppendText("Eroare: " + text);
                textBox13.AppendText(Environment.NewLine);

                string les = string.Join("\n", le);
                textBox13.AppendText(les);
                textBox13.AppendText(Environment.NewLine);
            }
            else
            {
                textBox13.AppendText("Nicio eroare: " + text);
                textBox13.AppendText(Environment.NewLine);
            }
        }

        void LogError(string text)
        {
            dynamic le = o.GetListaErori();
            if( le != null)
            {
                textBox13.AppendText("Eroare: " + text);
                textBox13.AppendText(Environment.NewLine);

                string les = string.Join("\n", le);
                textBox13.AppendText(les);
                textBox13.AppendText(Environment.NewLine);
            }  
            else
            {
                textBox13.AppendText("Nicio eroare, rezultat nul: " + text);
                textBox13.AppendText(Environment.NewLine);
            }
        }

       
         //varianta pt paradox nu are functia logon
        private void button1_Click(object sender, EventArgs e)
        {
            /*
           //   WMDocImpServer.WMDocImpObject o = new WMDocImpServer.WMDocImpObject();
           //    o.LogOn("GABI", "1");
           InitDocImpServer();
           string user = textBox6.Text;
           string pwd = textBox5.Text;
           int res = o.LogOn(user, pwd);
           if (res == 0)
               LogError("LogOn(" + user + ", " + pwd + ")");
           else
               LogOutput("LogOn(" + user + ", " + pwd + ")");
               */
        }


        private void button2_Click(object sender, EventArgs e)
        {
            InitDocImpServer();
            string firma = textBox4.Text;
            int res = o.SetNumeFirma(firma);
            if (res == 0)
                LogError("SetNumeFirma(" + firma + ")");
            else
                LogOutput("SetNumeFirma(" + firma + ")");
        }

        private void textBox4_TextChanged(object sender, EventArgs e)
        {

        }

        
        private void button3_Click(object sender, EventArgs e)
        {
            /*
            InitDocImpServer();
            string den = textBox3.Text;
            int res = o.SetDenSubunit(den);
            if (res == 0)
                LogError("SetDenSubunit(" + den + ")");
            else
                LogOutput("SetDenSubunit(" + den + ")");
                */
        }


        private void textBox3_TextChanged(object sender, EventArgs e)
        {

        }

        private void button4_Click(object sender, EventArgs e)
        {
            InitDocImpServer();
            int luna = Int32.Parse(textBox8.Text);
            int an = Int32.Parse(textBox7.Text);
            int res = o.SetLunaLucru(an, luna);
            if (res == 0)
                LogError("SetDenSubunit(" + an + ", " + luna + ")");
            else
                LogOutput("SetDenSubunit(" + an + ", " + luna + ")");
        }

        private void textBox8_TextChanged(object sender, EventArgs e)
        {

        }

        private void textBox7_TextChanged(object sender, EventArgs e)
        {

        }

        private void button5_Click(object sender, EventArgs e)
        {
            InitDocImpServer();
            string id = textBox9.Text;
            o.SetIDPartField(id);
            /*
            if (res == 0)
                LogError("SetDenSubunit(" + den + ")");
            else
            */
                LogOutput("SetIDPartField(" + id + ")");
        }

        private void textBox9_TextChanged(object sender, EventArgs e)
        {

        }

        private void button6_Click(object sender, EventArgs e)
        {
            InitDocImpServer();
            string id = textBox11.Text;
            o.SetIDArtField(id);
            /*
            if (res == 0)
                LogError("SetDenSubunit(" + den + ")");
            else
            */
                LogOutput("SetIDArtField(" + id + ")");
        }

        private void textBox11_TextChanged(object sender, EventArgs e)
        {
      
        }

        private void button7_Click(object sender, EventArgs e)
        {
            InitDocImpServer();
            dynamic arr = o.GetListaFirme();
            if (arr == null)
                LogError("GetListaFirme()");
            else
            {
                string s = string.Join(Environment.NewLine, arr);
                LogOutput("GetListaFirme()");
                LogOutput(s);
            }
           
        }

        private void button8_Click(object sender, EventArgs e)
        {
            int err = -1;
            InitDocImpServer();
            dynamic arr = o.GetListaSubunit(out err);
            if (arr == null)
                LogError("GetListaSubunit()");
            else
            {
                string s = string.Join(Environment.NewLine, arr);
                LogOutput("GetListaSubunitati()");
                LogOutput(s);
            }
        }

        private void button9_Click(object sender, EventArgs e)
        {
            int err = -1;
            InitDocImpServer();
            dynamic arr = o.GetComenziNefacturate(out err);
            if (arr == null)
                LogError("GetComenziNefacturate()");
            else
            {
                string s = string.Join(Environment.NewLine, arr);
                LogOutput("GetComenziNefacturate()");
                LogOutput(s);
            }
        }

        private void button10_Click(object sender, EventArgs e)
        {
            /*
            InitDocImpServer();
            int cod_cmd = Int32.Parse(textBox1.Text);
            int status = Int32.Parse(textBox2.Text);
            int res = o.SetStadiuWMSComanda(cod_cmd, status);
            if (res == 0)
                LogError("SetStadiuWMSComanda(" + cod_cmd + ", " + status + ")");
            else
                LogOutput("SetStadiuWMSComanda(" + cod_cmd + ", " + status + ")");
            */
        }

        private void textBox1_TextChanged(object sender, EventArgs e)
        {

        }

        private void Form1_Load(object sender, EventArgs e)
        {

        }

        private void textBox2_TextChanged(object sender, EventArgs e)
        {

        }

        private void button11_Click(object sender, EventArgs e)
        {
            /*
            InitDocImpServer();
            int err = -1;
            int cod_cmd = Int32.Parse(textBox10.Text);
            dynamic arr = o.GetLiniiComandaNefacturata(cod_cmd, out err);
            if (arr == null)
                LogError("GetLiniiComandaNefacturata("+ cod_cmd + ")");
            else
            {
                string s = string.Join(Environment.NewLine, arr);
                LogOutput("GetLiniiComandaNefacturata(" + cod_cmd + ")");
                LogOutput(s);
            }
            */
        }

        private void button12_Click(object sender, EventArgs e)
        {
            /*
            InitDocImpServer();
            string txt = textBox12.Text;
     //       string[] listaLinii = txt.Split(new[] { "\r\n", "\r", "\n" }, StringSplitOptions.None);
            string[] listaLinii = txt.Split(new[] { "\r\n", "\r", "\n" }, StringSplitOptions.RemoveEmptyEntries);

            int res = o.SetCantitatiLiniiComanda(listaLinii);
            if (res == 0)
                LogError("SetCantitatiLiniiComanda()");
            else
                LogOutput("SetCantitatiLiniiComanda()");
            */
        }

        private void textBox12_TextChanged(object sender, EventArgs e)
        {

        }

        private void textBox10_TextChanged(object sender, EventArgs e)
        {

        }

        private void button13_Click(object sender, EventArgs e)
        {
            /*
            InitDocImpServer();
            int err = -1;
            dynamic arr = o.GetCmdFurnNefacturate(out err);
            if (arr == null)
                LogError("GetLiniiComandaNefacturata()");
            else
            {
                string s = string.Join(Environment.NewLine, arr);
                LogOutput("GetLiniiComandaNefacturata()");
                LogOutput(s);
            }
            */
        }

        private void button14_Click(object sender, EventArgs e)
        {
            InitDocImpServer();
            int err = -1;
            dynamic arr = o.GetSolduriExt(out err);
            LogErrors("GetSolduriExt()");
            if (arr == null)
                LogError("GetSolduriExt()");
            else
            {
                string s = string.Join(Environment.NewLine, arr);
                LogOutput("GetSolduriExt()");
                LogOutput(s);
            }
        }
    }
}
